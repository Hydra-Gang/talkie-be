package models

import "github.com/google/uuid"

type UserGroup struct {
	BaseModel

	UserID  uuid.UUID `gorm:"type:uuid" json:"userId"`
	GroupID uuid.UUID `gorm:"type:uuid" json:"groupId"`

	User  *User  `json:"user"`
	Group *Group `json:"group"`
}
