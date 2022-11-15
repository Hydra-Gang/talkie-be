package creategroup

import (
	"github.com/Hydra-Gang/talkie-be/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func FindUserById(db *gorm.DB, id uuid.UUID) (models.User, error) {
	var user models.User
	err := db.Where("id = ?", id.String()).First(&user).Error

	return user, err
}

func CreateGroup(db *gorm.DB, name string) (models.Group, error) {
	group := models.Group{Name: name}
	err := db.Create(&group).Error

	return group, err
}

func CreateUserGroup(db *gorm.DB, groupId uuid.UUID, userId uuid.UUID) (models.UserGroup, error) {
	userGroup := models.UserGroup{
		GroupID: groupId,
		UserID: userId,
	}

	err := db.Create(&userGroup).Error
	return userGroup, err
}
