package kitchen

import (
	"context"
	"time"

	"github.com/dietdoctor/be-test/pkg/food/v1"
	"go.uber.org/zap"
)

//Handler is a kitchen handler
type Handler struct {
	log    *zap.SugaredLogger
	client food.KitchenServiceClient
}

//NewHandler creates a new kitchen handler
func NewHandler(log *zap.SugaredLogger, client food.KitchenServiceClient) *Handler {
	return &Handler{
		log:    log,
		client: client,
	}
}

//CreateKitchenOrder creates a kitchen order
func (h *Handler) CreateKitchenOrder(ctx context.Context, orderName string) (*food.KitchenOrder, error) {
	order, err := h.client.CreateKitchenOrder(ctx, &food.CreateKitchenOrderRequest{
		Kitchenorder: &food.KitchenOrder{
			Name:       orderName,
			CreateTime: time.Now().String(),
			Status:     food.KitchenOrder_NEW,
		},
	})
	if err != nil {
		h.log.Error(err)
		return nil, err
	}

	return order, nil
}

//GetKitchenOrder returns a single kitchen order
func (h *Handler) GetKitchenOrder(ctx context.Context, orderName string) (*food.KitchenOrder, error) {
	order, err := h.client.GetKitchenOrder(ctx, &food.GetKitchenOrderRequest{
		Name: orderName,
	})
	if err != nil {
		h.log.Error(err)
		return nil, err
	}

	return order, nil
}
