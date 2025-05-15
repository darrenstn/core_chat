package usecase

import (
	"core_chat/application/authentication/repository"
	"core_chat/application/authentication/service"
	"net/http"
)

type LoginUseCase struct {
	UserRepo     repository.UserRepository
	TokenService service.TokenService
	HashService  service.HashService
	HTTPService  service.HTTPService
}

func NewLoginUseCase(repo repository.UserRepository, ts service.TokenService, hs service.HashService, hsrv service.HTTPService) *LoginUseCase {
	return &LoginUseCase{UserRepo: repo, TokenService: ts, HashService: hs, HTTPService: hsrv}
}

func (uc *LoginUseCase) Execute(w http.ResponseWriter, username, password string) bool {
	user, err := uc.UserRepo.GetUserByUsername(username)
	if err != nil || !uc.HashService.CompareHash(user.Password, password) {
		return false
	}
	token, err := uc.TokenService.GenerateToken(user.UserName, user.UserType)
	if err != nil {
		return false
	}

	uc.HTTPService.SetAuthCookie(w, token)
	return true
}
