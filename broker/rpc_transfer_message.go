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
	go server.transferMessageToBroker(request.Message)
	// return a healthy empty response to the RPC caller
	return &pb.TransferMessageResponse{}, nil
}

// TODO: follow-id and status endpoint

// transferMessageToBroker forwards the message to the destination service using gRPC.
func (server *Server) transferMessageToBroker(message string) {
	// call the remote procedure
	// and create a new background context beacuse we don't need to cancel.
	// WARNING: this is a blocking call, so it will block the current goroutine.
	// WARNING: if the destination service is down or not responding, this will block
	// the current goroutine forever and it may cause memory leak.
	_, err := server.destinationClient.ProccessMessage(context.Background(), &pb.ProccessMessageRequest{Message: message})
	if err != nil {
		// if the RPC call returns an error, log it using the main logger
		log.Error().Str("ser_name", "broker_service").Err(err).Msg("failed to call remote procedure")
	}
}
