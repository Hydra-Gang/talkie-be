package creategroup

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UseCase(db *gorm.DB, dto DTO) {
	userId, _ := uuid.Parse(dto.UserID)
	targetUserId, _ := uuid.Parse(dto.TargetUserID)

	fmt.Println("User ID:", userId)
	fmt.Println("Target User ID:", targetUserId)
}
