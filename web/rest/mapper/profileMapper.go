package mapper

import (
	appdto "core_chat/application/person/dto"
	restdto "core_chat/web/rest/dto"
)

func ToProfileResult(res appdto.ProfileResult, pictureURL string, status int) restdto.ProfileResult {
	return restdto.ProfileResult{
		Status:  status,
		Message: res.Message,
		Data: restdto.ProfileData{
			Identifier:  res.Identifier,
			Name:        res.Name,
			Description: res.Description,
			PictureURL:  pictureURL,
		},
	}
}
