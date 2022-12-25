package receiver

import (
	"fmt"

	"github.com/SamMHD/simple-broker/pb"
	"github.com/gin-gonic/gin"
)

type forwardMessageRequest struct {
	Message string `json:"message" binding:"required"`
}

func (server *Server) forwardMessage(ctx *gin.Context) {
	var request forwardMessageRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, errorResponse(err))
		return
	}

	_, err := server.brokerClient.TransferMessage(ctx, &pb.TransferMessageRequest{Message: request.Message})
	fmt.Println(err)
	if err != nil {
		ctx.JSON(500, errorResponse(err))
		return
	}

	ctx.JSON(200, gin.H{"status": "ok"})
}
