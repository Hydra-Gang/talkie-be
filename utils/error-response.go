package utils

import (
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message    string
	StatusCode int
}

func SendErrorResponse(ctx *gin.Context, res ErrorResponse) {
	ctx.JSON(
		res.StatusCode,
		gin.H{
			"error": res.Message,
		})
}
