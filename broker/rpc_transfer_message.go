package broker

import (
	"context"

	"github.com/SamMHD/simple-broker/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) TransferMessage(ctx context.Context, request *pb.TransferMessageRequest) (*pb.TransferMessageResponse, error) {
	// TODO: Log message here
	// fmt.Println("Here", request.Message)

	_, err := server.destinationClient.ProccessMessage(ctx, &pb.ProccessMessageRequest{Message: request.Message})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to call remote procedure: %s", err)
	}
	return &pb.TransferMessageResponse{}, nil
}
