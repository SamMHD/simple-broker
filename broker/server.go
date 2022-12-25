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
	pb.UnimplementedBrokerServiceServer
	config util.Config
}

func NewServer(config util.Config) (*Server, error) {
	server := &Server{
		config: config,
	}

	return server, nil
}

func (server *Server) Start() {
	fmt.Println("Starting server...")
	lis, err := net.Listen("tcp", server.config.BrokerAddress)
	if err != nil {
		log.Fatal().Str("service", "DESTINATION").Msgf("failed to listen: %v", err)
	}

	logger := grpc.UnaryInterceptor(util.NewGrpcLoggerForService("destination_service"))
	s := grpc.NewServer(logger)
	pb.RegisterBrokerServiceServer(s, server)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatal().Str("service", "DESTINATION").Msgf("failed to serve: %v", err)
	}
}
