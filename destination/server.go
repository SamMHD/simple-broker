package destination

import (
	"fmt"
	"log"
	"net"

	"github.com/SamMHD/simple-broker/pb"
	"github.com/SamMHD/simple-broker/util"
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
	fmt.Println("Starting server...")
	lis, err := net.Listen("tcp", server.config.DestinationAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDestinationServiceServer(s, server)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
