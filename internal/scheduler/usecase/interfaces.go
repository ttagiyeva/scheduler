package usecase

import (
	"context"

	"github.com/dietdoctor/be-test/pkg/food/v1"
)

//go:generate mockery --with-expecter --name DroneService --testonly --case underscore --output ./mock --filename drone_service_mock.go --outpkg drone_service_mock --outpkg mock
type DroneService interface {
	CreateShipment(ctx context.Context, orderName string) (*food.Shipment, error)
	GetShipment(ctx context.Context, orderName string) (*food.Shipment, error)
}

//go:generate mockery --with-expecter --name OrderService --testonly --case underscore --output ./mock --filename order_service_mock.go --outpkg order_service_mock --outpkg mock
type OrderService interface {
	ListOrders(ctx context.Context, status food.Order_Status) ([]*food.Order, error)
	GetOrder(ctx context.Context, orderName string) (*food.Order, error)
	UpdateOrder(ctx context.Context, orderName string, status food.Order_Status) error
}

//go:generate mockery --with-expecter --name KitchenService --testonly --case underscore --output ./mock --filename kitchen_service_mock.go --outpkg kitchen_service_mock --outpkg mock
type KitchenService interface {
	CreateKitchenOrder(ctx context.Context, orderName string) (*food.KitchenOrder, error)
	GetKitchenOrder(ctx context.Context, orderName string) (*food.KitchenOrder, error)
}
