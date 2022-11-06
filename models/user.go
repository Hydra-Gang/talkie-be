package models

import "time"

type User struct {
	ID             string		`json:"id"`
	Email          string		`json:"email"`
	Password       string		`json:"password"`
	FullName       string		`json:"full_name"`
	Nickname       string		`json:"nickname"`
	DOB            string		`json:"dob"`
	Phone          string		`json:"phone"`
	ProfilePicture string		`json:"profile_picture"`
	RefreshToken   string		`json:"refresh_token"`
	UpdatedAt      time.Time 	`json:"updated_at"`
	CreatedAt      time.Time 	`json:"created_at"`
}
