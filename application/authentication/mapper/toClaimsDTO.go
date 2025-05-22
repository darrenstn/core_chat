package mapper

import (
	"core_chat/application/authentication/dto"
	"core_chat/application/authentication/model"
)

// ToClaimsDTO maps a Person object to a ClaimsDTO, adding additional info like ExpiresAt.
func ToClaimsDTO(p model.Person, expiresAt int64) dto.Claims {
	return dto.Claims{
		Identifier:     p.Identifier,
		Role:           p.Role,
		EmailValidated: p.EmailValidated,
		ExpiresAt:      expiresAt,
	}
}
