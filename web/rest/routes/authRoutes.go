package routes

import (
	"core_chat/application/authentication/usecase"
	"core_chat/web/rest"
	"core_chat/web/rest/dto"
	"core_chat/web/rest/util"

	// "encoding/json"

	"net/http"
)

type AuthHandler struct {
	LoginUC *usecase.LoginUseCase
}

func NewAuthHandler(loginUC *usecase.LoginUseCase) *AuthHandler {
	return &AuthHandler{LoginUC: loginUC}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// var req dto.LoginRequest
	// json.NewDecoder(r.Body).Decode(&req)
	identifier := r.FormValue("identifier")
	password := r.FormValue("password")

	req := dto.LoginRequest{
		Identifier: identifier,
		Password:   password,
	}

	result := h.LoginUC.Execute(req.Identifier, req.Password)

	if !result.Success {
		rest.SendResponse(w, 400, result.Message)
		return
	}

	util.SetAuthCookie(w, result.Token, result.Expiration)

	rest.SendResponse(w, 200, "Login success")
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	util.ClearAuthCookie(w)
	rest.SendResponse(w, 200, "Logout success")
}
