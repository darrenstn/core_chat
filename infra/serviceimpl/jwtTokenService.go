package serviceimpl

import (
	"core_chat/application/authentication/service"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtTokenService struct{}

func NewJWTTokenService() service.TokenService {
	return &jwtTokenService{}
}

func (s *jwtTokenService) GenerateToken(identifier string, role string) (string, error) {
	claims := jwt.MapClaims{
		"identifier": identifier,
		"role":       role,
		"exp":        time.Now().Add(5 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("test"))
}

func (s *jwtTokenService) ValidateToken(tokenStr string, requiredRole string) (bool, string, string) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte("test"), nil
	})
	if err != nil || !token.Valid {
		return false, "", ""
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		identifier, _ := claims["identifier"].(string)
		role, _ := claims["role"].(string)
		if role == requiredRole {
			return true, identifier, role
		}
	}
	return false, "", ""
}
