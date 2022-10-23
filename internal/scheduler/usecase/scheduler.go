package usecase

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/dietdoctor/be-test/pkg/food/v1"
	"github.com/ttagiyeva/scheduler/internal/drone"
	"github.com/ttagiyeva/scheduler/internal/kitchen"
	"github.com/ttagiyeva/scheduler/internal/order"
	"github.com/ttagiyeva/scheduler/internal/scheduler/domain"
	"github.com/ttagiyeva/scheduler/internal/scheduler/repository"
	"go.uber.org/zap"
)

const controllerInterval = 10

type Scheduler struct {
	drone   DroneService
	order   OrderService
	kitchen KitchenService
	repo    repository.Scheduler
	log     *zap.SugaredLogger
}

//New creates an Scheduler instance
func New(drone *drone.Handler, order *order.Handler, kitchen *kitchen.Handler, repo *repository.Firestore, log *zap.SugaredLogger) *Scheduler {
	return &Scheduler{
		drone:   drone,
		kitchen: kitchen,
		order:   order,
		repo:    repo,
		log:     log,
	}
}

//CreateKitchenOrders creates kitchen orders from orders which are not rejected or cancelled
func (s *Scheduler) CreateKitchenOrders(ctx context.Context) error {

	orders, err := s.order.ListOrders(ctx, food.Order_NEW)
	if err != nil {
		return err
	}

	orderNames := make([]string, 0, len(orders))

	for i := 0; i < len(orders); i++ {
		orderNames = append(orderNames, orders[i].Name)
	}

	if len(orderNames) != 0 { // empty slice cannot be used in firestore query

		partialOrders := make([]string, 0, 10)

		for {

			if len(orderNames) > 10 { // firestore query can have max 10 values
				partialOrders = orderNames[:10]
				orderNames = orderNames[10:]
			} else {
				partialOrders = orderNames
			}

			schedulers, err := s.repo.GetAll(ctx, "order_name", "in", partialOrders)
			if err != nil {
				return err
			}

			schedulerOrderNames := make([]string, 0, len(schedulers))

			for i := 0; i < len(schedulers); i++ {
				schedulerOrderNames = append(schedulerOrderNames, schedulers[i].OrderName)
			}

			newOrders := s.getDifference(orderNames, schedulerOrderNames)

			for _, order := range newOrders {
				kitchenOrder, err := s.kitchen.CreateKitchenOrder(ctx, order)
				if err != nil {
					return err
				}

				scheduler := &domain.Scheduler{
					DocumentId:  strings.Split(order, "/")[1],
					OrderName:   order,
					KitchenName: kitchenOrder.Name,
				}

				err = s.repo.Save(ctx, scheduler)
				if err != nil {
					return err
				}

			}

			if len(orderNames) <= 10 {
				return nil
			}
		}
	}

	return nil
}

//CreateShipmentOrders creates shipment orders from packaged kitchen orders which are not rejected or cancelled
func (s *Scheduler) CreateShipmentOrders(ctx context.Context) error {
	schedulers, err := s.repo.GetAll(ctx, "drone_name", "==", "")
	if err != nil {
		return err
	}

	for _, scheduler := range schedulers {

		order, err := s.order.GetOrder(ctx, scheduler.OrderName)
		if err != nil {
			return err
		}

		if order.Status == food.Order_REJECTED || order.Status == food.Order_CANCELLED {
			err = s.repo.Delete(ctx, scheduler.DocumentId)
			if err != nil {
				return err
			}

			s.log.Info("deleted rejected order from scheduler ", scheduler)

			continue
		}

		kitchenOrder, err := s.kitchen.GetKitchenOrder(ctx, scheduler.KitchenName)
		if err != nil {
			return err
		}

		if kitchenOrder.Status == food.KitchenOrder_PREPARATION {

			err = s.order.UpdateOrder(ctx, order.Name, food.Order_PREPARATION)
			if err != nil {
				return err
			}

			s.log.Info("updated order status to preparation ", scheduler)

		} else if kitchenOrder.Status == food.KitchenOrder_PACKAGED {

			shipment, err := s.drone.CreateShipment(ctx, scheduler.OrderName)
			if err != nil {
				return err
			}

			s.log.Info("created shipment order ", scheduler)

			scheduler.DroneName = shipment.Name

			err = s.repo.Update(ctx, scheduler)
			if err != nil {
				return err
			}

			s.log.Info("updated scheduler ", scheduler)
		}
	}

	return nil
}

//CompleteOrders completes orders which are delivered
func (s *Scheduler) CompleteOrders(ctx context.Context) error {

	schedulers, err := s.repo.GetAll(ctx, "drone_name", "!=", "")
	if err != nil {
		return err
	}

	for _, scheduler := range schedulers {

		order, err := s.order.GetOrder(ctx, scheduler.OrderName)
		if err != nil {
			return err
		}

		if order.Status == food.Order_REJECTED {
			err = s.repo.Delete(ctx, scheduler.DocumentId)
			if err != nil {
				return err
			}

			s.log.Info("deleted rejected order from scheduler ", scheduler)

			continue
		}

		shipment, err := s.drone.GetShipment(ctx, scheduler.DroneName)
		if err != nil {
			return err
		}

		if shipment.Status == food.Shipment_COLLECTED {

			err = s.order.UpdateOrder(ctx, order.Name, food.Order_IN_FLIGHT)
			if err != nil {
				return err
			}

			s.log.Info("updated order status to in_flight ", scheduler)

		} else if shipment.Status == food.Shipment_DELIVERED {

			err = s.repo.Delete(ctx, scheduler.DocumentId)
			if err != nil {
				return err
			}

			s.log.Info("order completed ", scheduler)

			err = s.order.UpdateOrder(ctx, order.Name, food.Order_DELIVERED)
			if err != nil {
				return err
			}

			s.log.Info("updated order status to delivered ", scheduler)
		}
	}

	return nil
}

// getDifference returns the difference between two slices
func (s *Scheduler) getDifference(first, second []string) []string {
	mb := make(map[string]struct{}, len(second))

	for _, x := range second {
		mb[x] = struct{}{}
	}

	var diff []string
	for _, x := range first {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}

	return diff
}

//Start starts the scheduler
func (s *Scheduler) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		var wg sync.WaitGroup

		wg.Add(3)

		go func() {
			s.CreateKitchenOrders(ctx)
			wg.Done()
		}()

		go func() {
			s.CreateShipmentOrders(ctx)
			wg.Done()
		}()

		go func() {
			s.CompleteOrders(ctx)
			wg.Done()
		}()

		wg.Wait()

		time.Sleep(time.Second * controllerInterval)
	}

}
