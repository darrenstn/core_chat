package dto

import "time"

type RegisterRequest struct {
	Identifier  string    `json:"identifier"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Description string    `json:"description"`
}
