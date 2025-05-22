package serviceimpl

import (
	"core_chat/application/authentication/dto"
	"core_chat/application/authentication/service"

	"github.com/golang-jwt/jwt"
)

type jwtTokenService struct{}

func NewJWTTokenService() service.TokenService {
	return &jwtTokenService{}
}

func (s *jwtTokenService) GenerateToken(claimsData dto.Claims) (string, error) {
	claims := jwt.MapClaims{
		"identifier":     claimsData.Identifier,
		"role":           claimsData.Role,
		"emailValidated": claimsData.EmailValidated,
		"exp":            claimsData.ExpiresAt,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("test"))
}

func (s *jwtTokenService) ValidateToken(tokenStr string, requiredRole string) (bool, string, string, bool) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte("test"), nil
	})
	if err != nil || !token.Valid {
		return false, "", "", false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		identifier, _ := claims["identifier"].(string)
		role, _ := claims["role"].(string)
		isEmailValidated, _ := claims["emailValidated"].(bool)
		if role == requiredRole {
			return true, identifier, role, isEmailValidated
		}
	}
	return false, "", "", false
}
