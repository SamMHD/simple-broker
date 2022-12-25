package destination

import (
	"context"
	"fmt"

	"github.com/SamMHD/simple-broker/pb"
)

func (server *Server) ProccessMessage(ctx context.Context, request *pb.ProccessMessageRequest) (*pb.ProccessMessageResponse, error) {
	fmt.Println("Received message: ", request.Message[:10], "...")
	return &pb.ProccessMessageResponse{}, nil
}
