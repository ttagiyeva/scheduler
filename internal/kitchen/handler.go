package kitchen

import (
	"context"
	"time"

	"github.com/dietdoctor/be-test/pkg/food/v1"
	"go.uber.org/zap"
)

//Handler kitchen service handler struct
type Handler struct {
	log    *zap.SugaredLogger
	Client food.KitchenServiceClient
}

//NewHandler creates a new kitchen handler instance
func NewHandler(log *zap.SugaredLogger, client food.KitchenServiceClient) *Handler {
	return &Handler{
		log:    log,
		Client: client}
}

func (h *Handler) CreateKitchenOrder(ctx context.Context, orderName string) (*food.KitchenOrder, error) {
	order, err := h.Client.CreateKitchenOrder(ctx, &food.CreateKitchenOrderRequest{
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

func (h *Handler) GetKitchenOrder(ctx context.Context, orderName string) (*food.KitchenOrder, error) {
	order, err := h.Client.GetKitchenOrder(ctx, &food.GetKitchenOrderRequest{
		Name: orderName,
	})
	if err != nil {
		h.log.Error(err)
		return nil, err
	}
	return order, nil
}
