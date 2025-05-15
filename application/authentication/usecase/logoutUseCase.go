package usecase

import (
	"core_chat/application/authentication/service"
	"net/http"
)

type LogoutUseCase struct {
	HTTPService service.HTTPService
}

func NewLogoutUseCase(h service.HTTPService) *LogoutUseCase {
	return &LogoutUseCase{HTTPService: h}
}

func (uc *LogoutUseCase) Execute(w http.ResponseWriter) {
	uc.HTTPService.ClearAuthCookie(w)
}
