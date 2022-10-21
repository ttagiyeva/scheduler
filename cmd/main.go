package main

import (
	"github.com/ttagiyeva/scheduler/internal/config"
	"github.com/ttagiyeva/scheduler/internal/drone"
	"github.com/ttagiyeva/scheduler/internal/kitchen"
	"github.com/ttagiyeva/scheduler/internal/log"
	"github.com/ttagiyeva/scheduler/internal/order"
	"github.com/ttagiyeva/scheduler/internal/scheduler/http"
	"github.com/ttagiyeva/scheduler/internal/scheduler/repository"
	"github.com/ttagiyeva/scheduler/internal/scheduler/usecase"
	"github.com/ttagiyeva/scheduler/internal/service"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.New,
			log.NewZapLogger,

			service.NewServer,

			drone.NewDroneClient,
			drone.NewHandler,

			kitchen.NewKitchenClient,
			kitchen.NewHandler,

			order.NewOrderClient,
			order.NewHandler,

			repository.NewFirestoreRepo,
			usecase.NewSchedulerUsecase,
			http.NewHandler,
		),
		fx.Invoke(service.RegisterRoutes),
	)

	app.Run()
}
