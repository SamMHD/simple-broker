package broker

import (
	"fmt"
	"net"

	"github.com/SamMHD/simple-broker/pb"
	"github.com/SamMHD/simple-broker/util"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	destinationClient pb.DestinationServiceClient
	config            util.Config

	pb.UnimplementedBrokerServiceServer
}

func NewServer(config util.Config) (*Server, error) {
	server := &Server{
		config: config,
	}

	return server, nil
}

func (server *Server) DialDestinationRPC() error {
	log.Info().Str("ser_name", "broker").Msg("Trying to dial Destination RPC...")

	conn, err := grpc.Dial(server.config.DestinationAddress, grpc.WithInsecure())
	if err != nil {
		return err
	}

	server.destinationClient = pb.NewDestinationServiceClient(conn)
	log.Info().Str("ser_name", "broker").Msg("Connection to Destination RPC established.")
	return nil
}

func (server *Server) StartGrpcServer() error {
	log.Info().Str("ser_name", "broker").Msg("Starting gRPC server...")

	lis, err := net.Listen("tcp", server.config.BrokerAddress)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	logger := grpc.UnaryInterceptor(util.NewGrpcLoggerForService("broker"))

	s := grpc.NewServer(logger)

	pb.RegisterBrokerServiceServer(s, server)
	reflection.Register(s)

	return s.Serve(lis)
}

func (server *Server) Start() {
	err := server.DialDestinationRPC()
	if err != nil {
		log.Fatal().Str("ser_name", "broker").Msgf("failed to dial destination: %s", err)
	}

	if err := server.StartGrpcServer(); err != nil {
		log.Fatal().Str("ser_name", "broker").Msgf("failed to serve: %v", err)
	}
}
