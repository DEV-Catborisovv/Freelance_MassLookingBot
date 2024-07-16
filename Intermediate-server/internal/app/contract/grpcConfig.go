package contract

import (
	"fmt"
	"net"

	"Freelance_MassLookingBot_Intermediate-server/internal/app/configs"
	"Freelance_MassLookingBot_Intermediate-server/internal/app/contract/gen/client_bot_proto"

	"google.golang.org/grpc"
)

type grpcServer struct {
	client_bot_proto.UnimplementedClientTelegramBotServer
}

func Init() error {
	config := configs.NewConfig()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GRPC.Port))
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	client_bot_proto.RegisterClientTelegramBotServer(server, &grpcServer{})

	if err := server.Serve(lis); err != nil {
		return err
	}
	return nil
}
