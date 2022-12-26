package receiver

import (
	"github.com/SamMHD/simple-broker/pb"
	"github.com/gin-gonic/gin"
)

// forwardMessageRequest is the request body for the "/forward" route.
type forwardMessageRequest struct {
	Message string `json:"message" binding:"required"`
}

// forwardMessage handles the "/forward" route.
// it forwards the message to the Broker over gRPC and translates the response to HTTP.
func (server *Server) forwardMessage(ctx *gin.Context) {
	// bind the request body to the request struct
	var request forwardMessageRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		// if the request body is invalid, return a 400 error
		ctx.JSON(400, errorResponse(err))
		return
	}

	// forwards the message the remote procedure
	_, err := server.brokerClient.TransferMessage(ctx, &pb.TransferMessageRequest{Message: request.Message})
	if err != nil {
		// if the RPC call returns an error, return a 500 error
		ctx.JSON(500, errorResponse(err))
		return
	}

	// return a 200 OK response
	ctx.JSON(200, gin.H{"status": "ok"})
}
