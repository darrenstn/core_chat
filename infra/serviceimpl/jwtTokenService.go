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

func (s *jwtTokenService) GenerateToken(identifier string, role string, emailValidated bool) (string, error) {
	claims := jwt.MapClaims{
		"identifier":     identifier,
		"role":           role,
		"emailValidated": emailValidated,
		"exp":            time.Now().Add(5 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("test"))
}

func (s *jwtTokenService) ValidateToken(tokenStr string, requiredRole string, emailValidated bool) (bool, string, string, bool) {
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
		if role == requiredRole && isEmailValidated == emailValidated {
			return true, identifier, role, isEmailValidated
		}
	}
	return false, "", "", false
}
