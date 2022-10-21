package service

import (
	"github.com/labstack/echo/v4"
	"github.com/ttagiyeva/scheduler/internal/scheduler/http"
)

//RegisterRoutes registers a path to a handler
func RegisterRoutes(e *echo.Echo, h *http.SchedulerHandler) error {
	e.GET("/health", healthcheck)
	e.POST("/kitchenorders", h.KitchenOrders)
	e.POST("/shipmentorders", h.ShipmentOrders)
	e.POST("/orders", h.Orders)
	return nil
}
