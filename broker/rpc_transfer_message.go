package broker

import (
	"context"

	"github.com/SamMHD/simple-broker/pb"

	"github.com/rs/zerolog/log"
)

// TransferMessage is the gRPC handler for the TransferMessage RPC.
func (server *Server) TransferMessage(ctx context.Context, request *pb.TransferMessageRequest) (*pb.TransferMessageResponse, error) {
	server.logger.Info().Str("ser_name", "broker_service").Str("passing_message", request.Message).Msg("Recieved Message")

	// after logging the message, it fires a new goroutine to forward the message to the destination service.
	go server.transferMessageToBroker(server.ctx, request.Message)
	// return a healthy empty response to the RPC caller
	return &pb.TransferMessageResponse{}, nil
}

// TODO: follow-id and status endpoint

// transferMessageToBroker forwards the message to the destination service using gRPC.
func (server *Server) transferMessageToBroker(ctx context.Context, message string) {

	// check if server context is done
	if ctx.Err() != nil {
		// if it is done, log it using the main logger
		log.Error().Str("ser_name", "broker_service").Msg("canceling message transfer. server context is done.")
	}

	// call the remote procedure using the gRPC client
	_, err := server.destinationClient.ProccessMessage(ctx, &pb.ProccessMessageRequest{Message: message})
	if err != nil {
		// if the RPC call returns an error, log it using the main logger
		log.Error().Str("ser_name", "broker_service").Err(err).Msg("failed to call remote procedure")
	}
}
