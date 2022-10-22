package order

import (
	"context"

	"github.com/dietdoctor/be-test/pkg/food/v1"
	"go.uber.org/zap"
	"google.golang.org/genproto/protobuf/field_mask"
)

//Handler order service handler struct
type Handler struct {
	log    *zap.SugaredLogger
	client food.OrderServiceClient
}

//NewHandler creates a new drone handler instance
func NewHandler(log *zap.SugaredLogger, client food.OrderServiceClient) *Handler {
	return &Handler{
		log:    log,
		client: client,
	}
}

func (h *Handler) ListOrders(ctx context.Context, status food.Order_Status) ([]*food.Order, error) {
	orders, err := h.client.ListOrders(ctx, &food.ListOrdersRequest{
		StatusFilter: status,
	})
	if err != nil {
		h.log.Error(err)
		return nil, err
	}

	return orders.Orders, nil
}

func (h *Handler) GetOrder(ctx context.Context, orderName string) (*food.Order, error) {
	order, err := h.client.GetOrder(ctx, &food.GetOrderRequest{
		Name: orderName,
	})
	if err != nil {
		h.log.Error(err)
		return nil, err
	}

	return order, nil
}

func (h *Handler) UpdateOrder(ctx context.Context, orderName string, status food.Order_Status) error {
	_, err := h.client.UpdateOrder(ctx, &food.UpdateOrderRequest{
		Order: &food.Order{
			Name:   orderName,
			Status: status,
		},
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"status"},
		},
	})
	if err != nil {
		h.log.Error(err)
		return err
	}

	return nil
}
