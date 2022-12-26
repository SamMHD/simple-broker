// package receiver implements the HTTP gateway for the Broker.
// it receives messages over HTTP and forwards them to the Broker using gRPC.
// HTTP requests are served by gin.
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
	config       util.Config            // configuration for the server
	router       *gin.Engine            // gin router
	brokerClient pb.BrokerServiceClient // gRPC client for the Broker
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config) (*Server, error) {
	// create a new server object
	server := &Server{
		config: config,
	}

	// connect to the Broker gRPC server
	err := server.connectBrokerClient()
	if err != nil {
		return nil, err
	}

	// set up gin router
	server.setupRouter()
	return server, nil
}

// setupRouter sets up the gin router and adds the routes.
// it also adds the logger and recovery middleware.
// path "/forward" is used to forward messages to the Broker.
func (server *Server) setupRouter() {
	// create a new gin router in order to avoid default logging middleware
	router := gin.New()
	// add the zerolog compatible logger middleware
	// here we use ginzerolog package to create a middleware which uses zerolog as the logger
	router.Use(ginzerolog.Logger("gin"))
	// add the recovery middleware to recover from panics
	// and return a 500 error
	router.Use(gin.Recovery())

	// add the route for forwarding messages
	router.POST("/forward", server.forwardMessage)

	// set the router to the server object
	server.router = router
}

// connectBrokerClient connects to the Broker gRPC server.
// it dial the server and creates a new gRPC client and sets it to the server object.
// it will also return an error if it fails to connect to the Broker or create a client.
func (server *Server) connectBrokerClient() error {
	log.Info().Str("ser_name", "receiver").Msg("Trying to dial Broker RPC...")

	// create a new gRPC connection to the Broker
	conn, err := grpc.Dial(server.config.BrokerAddress, grpc.WithInsecure())
	if err != nil {
		return err
	}

	// create a new gRPC client over the connection
	server.brokerClient = pb.NewBrokerServiceClient(conn)
	log.Info().Str("ser_name", "receiver").Msg("Connection to Broker RPC established.")

	// return nil error
	return nil
}

// Start runs the Gin HTTP server
// it will return an error if the server fails to start, otherwise it will block the thread
// and wait for the server to stop. in case of peacefull stop, it will return nil.
func (server *Server) Start() error {
	log.Info().Str("ser_name", "receiver").Msg("Starting gin server...")
	return server.router.Run(server.config.ReceiverAddress)
}

// errorResponse is a helper function to create a gin.H map with the error key and the error message.
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
