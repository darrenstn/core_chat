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

func (s *jwtTokenService) GenerateToken(username string, userType int) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"type":     userType,
		"exp":      time.Now().Add(5 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("test"))
}

func (s *jwtTokenService) ValidateToken(tokenStr string, requiredType int) (bool, string, int) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte("test"), nil
	})
	if err != nil || !token.Valid {
		return false, "", -1
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		username, _ := claims["username"].(string)
		typeFloat, _ := claims["type"].(float64)
		userType := int(typeFloat)
		if userType == requiredType {
			return true, username, userType
		}
	}
	return false, "", -1
}
