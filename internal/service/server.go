package service

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ttagiyeva/scheduler/internal/config"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

//NewServer creates an echo instance
func NewServer(lc fx.Lifecycle, config *config.Config, logger *zap.SugaredLogger) (*echo.Echo, error) {
	engine := echo.New()
	engine.Use(middleware.Logger()) //?
	engine.Use(middleware.Recover())

	errCh := make(chan error)
	succCh := make(chan int)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				addr := fmt.Sprintf(":%d", config.HttpServerConfig.Port)
				listener, err := net.Listen("tcp", addr)
				if err != nil {
					logger.Error(err)
					errCh <- err
				}
				engine.Listener = listener

				succCh <- 0

				if err := engine.Start(addr); err != nil && err != http.ErrServerClosed {
					logger.Error(err)
					errCh <- err
				}
			}()

			select {
			case <-succCh:
				return nil
			case e := <-errCh:
				return e
			case <-ctx.Done():
				return ctx.Err()
			}
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("shutting down the server gracefully")
			return engine.Shutdown(ctx)
		},
	})

	return engine, nil
}
