package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ttagiyeva/scheduler/internal/rest"
	"github.com/ttagiyeva/scheduler/internal/scheduler/usecase"
)

type SchedulerHandler struct {
	usecase *usecase.SchedulerUsecase
}

func NewHandler(usecase *usecase.SchedulerUsecase) *SchedulerHandler {
	return &SchedulerHandler{
		usecase: usecase,
	}
}

func (s *SchedulerHandler) KitchenOrders(c echo.Context) error {
	ctx := c.Request().Context()
	err := s.usecase.CreateKitchenOrders(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, rest.InternalServerErrorResponse())
	}
	return c.JSON(http.StatusOK, nil) //?
}

func (s *SchedulerHandler) ShipmentOrders(c echo.Context) error {
	ctx := c.Request().Context()
	err := s.usecase.CreateShipmentOrders(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, rest.InternalServerErrorResponse())
	}
	return c.JSON(http.StatusOK, nil) //?
}

func (s *SchedulerHandler) Orders(c echo.Context) error {
	ctx := c.Request().Context()
	err := s.usecase.CompleteOrders(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, rest.InternalServerErrorResponse())
	}
	return c.JSON(http.StatusOK, nil) //?
}
