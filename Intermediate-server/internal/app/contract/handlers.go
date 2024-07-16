package contract

import (
	"Freelance_MassLookingBot_Intermediate-server/internal/app/contract/gen/client_bot_proto"
	"context"
)

func (s *grpcServer) GetPhoneNumber(ctx context.Context, req *client_bot_proto.GetPhoneNumberRequest) (*client_bot_proto.GetPhoneNumberResponse, error) {
	return &client_bot_proto.GetPhoneNumberResponse{
		//
	}, nil
}

func (s *grpcServer) GetCode(ctx context.Context, req *client_bot_proto.GetCodeRequest) (*client_bot_proto.GetCodeResponse, error) {
	return &client_bot_proto.GetCodeResponse{
		//
	}, nil
}
