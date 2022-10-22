package kitchen

import (
	kitchenClient "github.com/dietdoctor/be-test/pkg/food/v1"
	"github.com/ttagiyeva/scheduler/internal/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// NewClient is new grpc client for Kitchen service
func NewClient(log *zap.SugaredLogger, cfg *config.Config) kitchenClient.KitchenServiceClient {
	creds, err := credentials.NewClientTLSFromFile("server.crt", "")
	if err != nil {
		log.Fatalf("could not process the credentials: %v", err)
	}

	conn, err := grpc.Dial(cfg.GrpcServerConfig.Port, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}

	kitchenClient := kitchenClient.NewKitchenServiceClient(conn)

	return kitchenClient
}
