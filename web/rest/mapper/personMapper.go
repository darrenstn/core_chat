package mapper

import (
	appdto "core_chat/application/person/dto"
	restdto "core_chat/web/rest/dto"
)

func ToRegisterRequest(req restdto.RegisterRequest, picturePath string) appdto.RegisterRequest {
	return appdto.RegisterRequest{
		Identifier:  req.Identifier,
		Email:       req.Email,
		Password:    req.Password,
		Name:        req.Name,
		DateOfBirth: req.DateOfBirth,
		Description: req.Description,
		PicturePath: picturePath,
	}
}
