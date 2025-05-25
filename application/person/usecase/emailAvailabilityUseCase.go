package usecase

import (
	"core_chat/application/person/repository"
	"core_chat/application/person/service"
)

type EmailAvailabilityUseCase struct {
	Repo      repository.PersonRepository
	Validator service.ValidatorService
}

func NewEmailAvailabilityUseCase(repo repository.PersonRepository, vs service.ValidatorService) *EmailAvailabilityUseCase {
	return &EmailAvailabilityUseCase{Repo: repo, Validator: vs}
}

func (uc *EmailAvailabilityUseCase) Execute(email string) bool {
	return isEmailAvailable(email, uc.Repo, uc.Validator)
}
