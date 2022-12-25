package broker

import (
	"context"
	"fmt"

	"github.com/SamMHD/simple-broker/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) TransferMessage(ctx context.Context, request *pb.TransferMessageRequest) (*pb.TransferMessageResponse, error) {
	fmt.Println("transfering message: ", request.Message[:10], "...")

	// create a grpc client
	conn, err := grpc.Dial(server.config.DestinationAddress, grpc.WithInsecure())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to dial destination: %s", err)
	}
	defer conn.Close()

	// TODO: Log message here
	// fmt.Println("Here", request.Message)

	client := pb.NewDestinationServiceClient(conn)
	_, err = client.ProccessMessage(ctx, &pb.ProccessMessageRequest{Message: request.Message})
	if err != nil {
		fmt.Println("Error", err)
		return nil, status.Errorf(codes.Internal, "failed to call remote procedure: %s", err)
	}

	return &pb.TransferMessageResponse{}, nil
}
