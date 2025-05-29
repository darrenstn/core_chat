package usecase

import (
	"core_chat/application/authentication/dto"
	"core_chat/application/authentication/mapper"
	"core_chat/application/authentication/model"
	"core_chat/application/authentication/service"
	"core_chat/config"
	"time"
)

func generateAuthResult(person model.Person, tokenService service.TokenService, successMessage string) dto.AuthResult {
	expTime := time.Now().Add(config.AccessTokenExpiration)
	claims := mapper.ToClaimsDTO(person, expTime.Unix())

	token, err := tokenService.GenerateToken(claims)
	if err != nil {
		return dto.AuthResult{Success: false, Message: "Token generation failed"}
	}

	return dto.AuthResult{
		Success:    true,
		Token:      token,
		Message:    successMessage,
		Expiration: expTime,
	}
}
