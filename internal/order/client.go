package order

import (
	orderClient "github.com/dietdoctor/be-test/pkg/food/v1"
	"github.com/ttagiyeva/scheduler/internal/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// NewOrderClient is new grpc client for Order service
func NewOrderClient(log *zap.SugaredLogger, cfg *config.Config) orderClient.OrderServiceClient {
	conn, err := grpc.Dial(cfg.GrpcServerConfig.Port, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		panic(err)
	}
	orderClient := orderClient.NewOrderServiceClient(conn)

	return orderClient
}
