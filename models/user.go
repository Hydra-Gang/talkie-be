package models

import (
	"time"
)

type User struct {
	BaseModel
	Email          string    `json:"email"`
	Password       string    `gorm:"not null"`
	FullName       string    `gorm:"not null" json:"fullName"`
	Nickname       string    `json:"nickName"`
	DOB            time.Time `gorm:"not null" json:"dob"`
	Phone          string    `gorm:"not null" json:"phone"`
	ProfilePicture string    `json:"profilePicture"`
	RefreshToken   string    `json:"refreshToken"`
}
