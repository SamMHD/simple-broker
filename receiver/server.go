package receiver

import (
	"github.com/SamMHD/simple-broker/util"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for the receiver
type Server struct {
	config util.Config
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config) (*Server, error) {
	server := &Server{
		config: config,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/forward", server.forwardMessage)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start() error {
	return server.router.Run(server.config.ReceiverAddress)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
