package entity

import (
	"time"
)

type Person struct {
	Identifier  string
	Password    string
	Role        string
	Name        string
	Email       string
	DateOfBirth time.Time
	Description string
	PicturePath string
}
