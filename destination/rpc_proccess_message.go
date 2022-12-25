package destination

import (
	"context"

	"github.com/SamMHD/simple-broker/pb"
)

// var int recieveMessageCount = 0

func (server *Server) ProccessMessage(ctx context.Context, request *pb.ProccessMessageRequest) (*pb.ProccessMessageResponse, error) {
	return &pb.ProccessMessageResponse{}, nil
}
