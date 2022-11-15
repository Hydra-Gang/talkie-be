package creategroup

import (
	"net/http"

	"github.com/Hydra-Gang/talkie-be/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DTO struct {
	UserID       string `json:"userId" binding:"uuid,required"`
	TargetUserID string `json:"targetUserId" binding:"uuid,required"`
}

type Controller struct {
	db *gorm.DB
}

func New(db *gorm.DB) gin.HandlerFunc {
	return Controller{db: db}.handler
}

func (c Controller) handler(ctx *gin.Context) {
	var dto DTO

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		utils.SendErrorResponse(ctx, utils.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	UseCase(c.db, dto)
}
