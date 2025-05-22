package service

import (
	"core_chat/application/authentication/dto"
)

type TokenService interface {
	GenerateToken(claims dto.Claims) (string, error)
	ValidateToken(tokenString string, role string) (bool, string, string, bool)
}
