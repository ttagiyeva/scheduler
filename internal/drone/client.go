package drone

import (
	"crypto/tls"

	droneClient "github.com/dietdoctor/be-test/pkg/food/v1"
	"github.com/ttagiyeva/scheduler/internal/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// NewClient is new grpc client for Drone service
func NewClient(log *zap.SugaredLogger, cfg *config.Config) droneClient.DroneServiceClient {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := grpc.Dial(cfg.GrpcServerConfig.Port, grpc.WithTransportCredentials(credentials.NewTLS(config)))
	if err != nil {
		log.Fatal(err)
	}

	droneClient := droneClient.NewDroneServiceClient(conn)

	return droneClient
}
