package usecase

import (
	"core_chat/application/person/repository"
	"core_chat/application/person/service"
)

type IdentifierAvailabilityUseCase struct {
	Repo      repository.PersonRepository
	Validator service.ValidatorService
}

func NewIdentifierAvailabilityUseCase(repo repository.PersonRepository, vs service.ValidatorService) *IdentifierAvailabilityUseCase {
	return &IdentifierAvailabilityUseCase{Repo: repo, Validator: vs}
}

func (uc *IdentifierAvailabilityUseCase) Execute(identifier string) bool {
	return isIdentifierAvailable(identifier, uc.Repo, uc.Validator)
}
