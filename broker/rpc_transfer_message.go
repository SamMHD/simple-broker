package broker

import (
	"context"

	"github.com/SamMHD/simple-broker/pb"

	"github.com/rs/zerolog/log"
)

func (server *Server) TransferMessage(ctx context.Context, request *pb.TransferMessageRequest) (*pb.TransferMessageResponse, error) {
	// TODO: Log message here
	log.Info().Str("ser_name", "broker_service").Str("passing_message", request.Message).Msg("Recieved Message")

	// if request.Message[0] == 'A' {
	// 	fmt.Println("Here", request.Message[:10])
	// }
	go server.TransferMessageToBroker(request.Message)
	return &pb.TransferMessageResponse{}, nil
}

// TODO: follow-id and status endpoint

func (server *Server) TransferMessageToBroker(message string) {
	_, err := server.destinationClient.ProccessMessage(context.Background(), &pb.ProccessMessageRequest{Message: message})
	if err != nil {
		// return nil, status.Errorf(codes.Internal, "failed to call remote procedure: %s", err)
		log.Error().Str("ser_name", "broker_service").Msgf("failed to call remote procedure: %s", err)
	}
}
