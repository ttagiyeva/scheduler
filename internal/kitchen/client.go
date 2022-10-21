package kitchen

import (
	kitchenClient "github.com/dietdoctor/be-test/pkg/food/v1"
	"github.com/ttagiyeva/scheduler/internal/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// NewKitchenClient is new grpc client for Kitchen service
func NewKitchenClient(log *zap.SugaredLogger, cfg *config.Config) kitchenClient.KitchenServiceClient {
	conn, err := grpc.Dial(cfg.GrpcServerConfig.Port, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		panic(err)
	}
	kitchenClient := kitchenClient.NewKitchenServiceClient(conn)

	return kitchenClient
}
