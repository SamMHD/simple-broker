package receiver

import "github.com/gin-gonic/gin"

type forwardMessageRequest struct {
	Message string `json:"message" binding:"required"`
}

func (s *Server) forwardMessage(ctx *gin.Context) {
	var request forwardMessageRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, errorResponse(err))
		return
	}

	// TODO: send message over gRPC
}
