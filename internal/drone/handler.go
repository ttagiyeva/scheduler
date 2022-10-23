package drone

import (
	"context"
	"time"

	"github.com/dietdoctor/be-test/pkg/food/v1"
	"go.uber.org/zap"
)

//Handler is a drone handler
type Handler struct {
	log    *zap.SugaredLogger
	client food.DroneServiceClient
}

//NewHandler creates a new drone handler
func NewHandler(log *zap.SugaredLogger, client food.DroneServiceClient) *Handler {
	return &Handler{
		log:    log,
		client: client,
	}
}

//CreateShipment creates a shipment
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

//UpdateShipment updates a shipment
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
