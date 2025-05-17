package routes

import (
	"core_chat/application/authentication/usecase"
	"core_chat/web/rest"
	"core_chat/web/rest/dto"

	// "encoding/json"

	"net/http"
)

type AuthHandler struct {
	LoginUC  *usecase.LoginUseCase
	LogoutUC *usecase.LogoutUseCase
}

func NewAuthHandler(loginUC *usecase.LoginUseCase, logoutUC *usecase.LogoutUseCase) *AuthHandler {
	return &AuthHandler{LoginUC: loginUC, LogoutUC: logoutUC}
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

	if ok := h.LoginUC.Execute(w, req.Identifier, req.Password); !ok {
		rest.SendResponse(w, 400, "Invalid credentials")
		return
	}
	rest.SendResponse(w, 200, "Login success")
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	h.LogoutUC.Execute(w)
	rest.SendResponse(w, 200, "Logout success")
}
