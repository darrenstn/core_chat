package usecase

import (
	"core_chat/application/person/repository"
	"core_chat/application/person/service"
)

func isIdentifierAvailable(identifier string, repo repository.PersonRepository, validator service.ValidatorService) bool {
	if !validator.IsIdentifierValid(identifier) {
		return false
	}
	return !repo.ExistsByIdentifier(identifier)
}

func isEmailAvailable(email string, repo repository.PersonRepository, validator service.ValidatorService) bool {
	if !validator.IsEmailValid(email) {
		return false
	}
	return !repo.ExistsByEmail(email)
}
