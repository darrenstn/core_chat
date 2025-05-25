package dto

import (
	"time"
)

type RegisterRequest struct {
	Identifier  string
	Password    string
	Name        string
	Email       string
	DateOfBirth time.Time
	Description string
	PicturePath string
}
