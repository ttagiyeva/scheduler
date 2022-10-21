package usecase

import (
	"context"

	"github.com/dietdoctor/be-test/pkg/food/v1"
	"github.com/ttagiyeva/scheduler/internal/drone"
	"github.com/ttagiyeva/scheduler/internal/kitchen"
	"github.com/ttagiyeva/scheduler/internal/order"
	"github.com/ttagiyeva/scheduler/internal/scheduler/domain"
	"github.com/ttagiyeva/scheduler/internal/scheduler/repository"
)

type SchedulerUsecase struct {
	drone   *drone.Handler
	order   *order.Handler
	kitchen *kitchen.Handler
	repo    repository.SchedulerRepo
}

//NewSchedulerUsecase creates an SchedulerUsecase instance
func NewSchedulerUsecase(drone *drone.Handler, order *order.Handler, kitchen *kitchen.Handler, repo *repository.FirestoreRepo) *SchedulerUsecase {
	return &SchedulerUsecase{
		drone:   drone,
		kitchen: kitchen,
		order:   order,
		repo:    repo,
	}
}

//CreateKitchenOrders creates kitchen order to orders that are new
func (s *SchedulerUsecase) CreateKitchenOrders(ctx context.Context) error {

	orders, err := s.order.ListOrders(ctx, food.Order_NEW)
	if err != nil {
		return err
	}

	for _, order := range orders {
		kitchenOrder, err := s.kitchen.CreateKitchenOrder(ctx, order.Name)
		if err != nil {
			return err
		}

		scheduler := &domain.Scheduler{
			OrderName:   order.Name,
			KitchenName: kitchenOrder.Name,
		}

		err = s.repo.Save(ctx, scheduler)
		if err != nil {
			return err
		}
	}

	return nil
}

//CreateShipmentOrders creates shipment order to kitchen orders which are packaged
func (s *SchedulerUsecase) CreateShipmentOrders(ctx context.Context) error {
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
			err = s.repo.Delete(ctx, scheduler.OrderName)
			if err != nil {
				return err
			}
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

		} else if kitchenOrder.Status == food.KitchenOrder_PACKAGED {

			shipment, err := s.drone.CreateShipment(ctx, scheduler.OrderName)
			if err != nil {
				return err
			}

			scheduler.DroneName = shipment.Name

			err = s.repo.Update(ctx, scheduler)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

//CompleteOrders changes order status depend on shipment status
func (s *SchedulerUsecase) CompleteOrders(ctx context.Context) error {

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
			err = s.repo.Delete(ctx, scheduler.OrderName)
			if err != nil {
				return err
			}
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

		} else if shipment.Status == food.Shipment_DELIVERED {

			err = s.repo.Delete(ctx, scheduler.OrderName)
			if err != nil {
				return err
			}

			err = s.order.UpdateOrder(ctx, order.Name, food.Order_DELIVERED)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
