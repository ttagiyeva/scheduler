package drone

import (
	"context"
	"time"

	"github.com/dietdoctor/be-test/pkg/food/v1"
	"go.uber.org/zap"
)

//Handler is drone service handler struct
type Handler struct {
	log    *zap.SugaredLogger
	client food.DroneServiceClient
}

//NewHandler creates a new drone handler instance
func NewHandler(log *zap.SugaredLogger, client food.DroneServiceClient) *Handler {
	return &Handler{
		log:    log,
		client: client,
	}
}

func (h *Handler) CreateShipment(ctx context.Context, orderName string) (*food.Shipment, error) {
	shipment, err := h.client.CreateShipment(ctx, &food.CreateShipmentRequest{
		Shipment: &food.Shipment{
			Name:       orderName,
			CreateTime: time.Now().String(),
			Status:     food.Shipment_NEW,
		},
	})
	if err != nil {
		h.log.Error(err)
		return nil, err
	}
	return shipment, nil
}

func (h *Handler) GetShipment(ctx context.Context, orderName string) (*food.Shipment, error) {
	shipment, err := h.client.GetShipment(ctx, &food.GetShipmentRequest{
		Name: orderName,
	})
	if err != nil {
		h.log.Error(err)
		return nil, err
	}
	return shipment, nil
}
