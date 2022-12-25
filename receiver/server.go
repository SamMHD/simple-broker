package receiver

import (
	"github.com/SamMHD/simple-broker/pb"
	"github.com/SamMHD/simple-broker/util"
	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// Server serves HTTP requests for the receiver
type Server struct {
	config       util.Config
	router       *gin.Engine
	brokerClient pb.BrokerServiceClient
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config) (*Server, error) {
	server := &Server{
		config: config,
	}

	err := server.connectBrokerClient()
	if err != nil {
		return nil, err
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.New()
	router.Use(ginzerolog.Logger("gin"))
	router.Use(gin.Recovery())

	router.POST("/forward", server.forwardMessage)

	server.router = router
}

func (server *Server) connectBrokerClient() error {
	log.Info().Str("ser_name", "receiver").Msg("Trying to dial Broker RPC...")

	conn, err := grpc.Dial(server.config.BrokerAddress, grpc.WithInsecure())
	if err != nil {
		return err
	}

	server.brokerClient = pb.NewBrokerServiceClient(conn)
	log.Info().Str("ser_name", "receiver").Msg("Connection to Broker RPC established.")
	return nil
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start() error {
	log.Info().Str("ser_name", "receiver").Msg("Starting gin server...")
	return server.router.Run(server.config.ReceiverAddress)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
