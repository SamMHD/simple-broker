package destination

import (
	"net"

	"github.com/SamMHD/simple-broker/pb"
	"github.com/SamMHD/simple-broker/util"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedDestinationServiceServer
	config util.Config
}

func NewServer(config util.Config) (*Server, error) {
	server := &Server{
		config: config,
	}

	return server, nil
}

func (server *Server) Start() {
	log.Info().Str("ser_name", "des_service").Msg("starting server...")
	lis, err := net.Listen("tcp", server.config.DestinationAddress)
	if err != nil {
		log.Fatal().Str("ser_name", "des_service").Msgf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDestinationServiceServer(s, server)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatal().Str("ser_name", "des_service").Msgf("failed to listen: %v", err)
	}
}
