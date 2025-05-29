package usecase

import (
	"core_chat/application/person/dto"
	"core_chat/application/person/repository"
	"core_chat/application/person/service"
)

type GetProfileImageUseCase struct {
	Repo             repository.PersonRepository
	ImageService     service.ImageService
	DefaultImagePath string
}

func NewGetProfileImageUseCase(repo repository.PersonRepository, imageService service.ImageService, defaultPath string) *GetProfileImageUseCase {
	return &GetProfileImageUseCase{
		Repo:             repo,
		ImageService:     imageService,
		DefaultImagePath: defaultPath,
	}
}

func (uc *GetProfileImageUseCase) Execute(identifier string) dto.ProfileImageResult {
	profile, err := uc.Repo.FindProfileByIdentifier(identifier)
	if err != nil {
		return dto.ProfileImageResult{
			Success: false,
			Message: "Person not found",
		}
	}

	imagePath := uc.ImageService.GetProfileImagePathOrDefault(profile.PicturePath, uc.DefaultImagePath)

	return dto.ProfileImageResult{
		Success:     true,
		PicturePath: imagePath,
	}
}
