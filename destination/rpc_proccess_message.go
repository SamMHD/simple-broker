package destination

import (
	"context"

	"github.com/SamMHD/simple-broker/pb"
	"google.golang.org/protobuf/proto"

	"github.com/rs/zerolog/log"
)

var recievedMessagesCount int32
var recievedMessagesTotalSize int64

func (server *Server) ProccessMessage(ctx context.Context, request *pb.ProccessMessageRequest) (*pb.ProccessMessageResponse, error) {
	messageSize := int64(proto.Size(request))
	recievedMessagesTotalSize += messageSize
	recievedMessagesCount++

	log.Info().Str("ser_name", "des_service").
		Int32("TotalMessages", recievedMessagesCount).
		Int64("TotalSize", recievedMessagesTotalSize).
		Int64("MessageSize", messageSize).
		Msg("Recieved Message")

	return &pb.ProccessMessageResponse{}, nil
}
