package usecase

import (
	"core_chat/application/person/dto"
	"core_chat/application/person/entity"
	"core_chat/application/person/repository"
	"core_chat/application/person/service"
	"fmt"
)

type RegisterPersonUseCase struct {
	Repo        repository.PersonRepository
	HashService service.HashService
	Antivirus   service.AntivirusService
	Validator   service.ValidatorService
	ImageSvc    service.ImageService
}

func NewRegisterPersonUseCase(repo repository.PersonRepository, hs service.HashService, avs service.AntivirusService, vs service.ValidatorService, is service.ImageService) *RegisterPersonUseCase {
	return &RegisterPersonUseCase{Repo: repo, HashService: hs, Antivirus: avs, Validator: vs, ImageSvc: is}
}

func (uc *RegisterPersonUseCase) Execute(req dto.RegisterRequest) dto.RegisterResult {
	if !isIdentifierAvailable(req.Identifier, uc.Repo, uc.Validator) {
		return dto.RegisterResult{Success: false, Message: "Identifier already in use or not valid"}
	}

	if !isEmailAvailable(req.Email, uc.Repo, uc.Validator) {
		return dto.RegisterResult{Success: false, Message: "Email already in use or not valid"}
	}

	if !uc.Validator.IsPasswordValid(req.Password) {
		return dto.RegisterResult{Success: false, Message: "Password not valid"}
	}

	newPicturePath := ""

	if req.PicturePath != "" {
		if err := uc.Antivirus.ScanImage(req.PicturePath); err != nil {
			fmt.Println("Virus scan error:", err)
			return dto.RegisterResult{Success: false, Message: "Profile picture failed virus scan"}
		}

		var err error

		if newPicturePath, err = uc.ImageSvc.ResizeImage(req.PicturePath, "image/profile", 300, 300); err != nil {
			fmt.Println("Resize image error:", err)
			return dto.RegisterResult{Success: false, Message: "Failed to resize profile picture"}
		}
	}

	hashedPassword, err := uc.HashService.HashPassword(req.Password)

	if err != nil {
		return dto.RegisterResult{Success: false, Message: "Error hashing password"}
	}

	person := entity.Person{
		Identifier:  req.Identifier,
		Password:    hashedPassword,
		Role:        "user",
		Name:        req.Name,
		Email:       req.Email,
		DateOfBirth: req.DateOfBirth,
		Description: req.Description,
		PicturePath: newPicturePath,
	}

	if req.PicturePath == "" {
		person.PicturePath = `upload\image\profile\default.jpg`
	}

	err = uc.Repo.SavePerson(person)
	if err != nil {
		return dto.RegisterResult{Success: false, Message: "Failed to register person"}
	}

	return dto.RegisterResult{Success: true, Message: "Registration successful"}
}
