package main

import (
	"context"

	"github.com/ttagiyeva/scheduler/internal/config"
	"github.com/ttagiyeva/scheduler/internal/drone"
	"github.com/ttagiyeva/scheduler/internal/kitchen"
	"github.com/ttagiyeva/scheduler/internal/log"
	"github.com/ttagiyeva/scheduler/internal/order"
	"github.com/ttagiyeva/scheduler/internal/scheduler/repository"
	"github.com/ttagiyeva/scheduler/internal/scheduler/usecase"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.WithLogger(func(l *zap.SugaredLogger) fxevent.Logger {
			return &fxevent.ZapLogger{
				Logger: l.Desugar(),
			}
		}),

		fx.Provide(
			config.New,
			log.NewZapLogger,

			drone.NewClient,
			drone.NewHandler,

			kitchen.NewClient,
			kitchen.NewHandler,

			order.NewClient,
			order.NewHandler,

			repository.New,
			usecase.New,
		),
		fx.Invoke(func(uc *usecase.Scheduler) {
			go uc.Start(context.Background())
		}),
	).Run()
}
