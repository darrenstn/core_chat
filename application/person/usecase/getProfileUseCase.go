package usecase

import (
	"core_chat/application/person/dto"
	"core_chat/application/person/repository"
)

type GetProfileUseCase struct {
	Repo repository.PersonRepository
}

func NewGetProfileUseCase(repo repository.PersonRepository) *GetProfileUseCase {
	return &GetProfileUseCase{Repo: repo}
}

func (uc *GetProfileUseCase) Execute(identifier string) dto.ProfileResult {
	profile, err := uc.Repo.FindProfileByIdentifier(identifier)
	if err != nil {
		return dto.ProfileResult{Success: false, Message: "Person not found"}
	}

	return dto.ProfileResult{
		Success:     true,
		Message:     "Successfully retrieved",
		Identifier:  profile.Identifier,
		Name:        profile.Name,
		Description: profile.Description,
		PicturePath: profile.PicturePath,
	}
}
