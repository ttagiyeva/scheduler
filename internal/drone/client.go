package drone

import (
	droneClient "github.com/dietdoctor/be-test/pkg/food/v1"
	"github.com/ttagiyeva/scheduler/internal/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// NewDroneClient is new grpc client for Drone service
func NewDroneClient(log *zap.SugaredLogger, cfg *config.Config) droneClient.DroneServiceClient {
	conn, err := grpc.Dial(cfg.GrpcServerConfig.Port, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		panic(err)
	}
	droneClient := droneClient.NewDroneServiceClient(conn)

	return droneClient
}
