package routes

import (
	"core_chat/application/authentication/usecase"
	"core_chat/web/rest"
	"core_chat/web/rest/dto"
	restutil "core_chat/web/rest/util"
	webutil "core_chat/web/util"

	// "encoding/json"

	"net/http"
)

type AuthHandler struct {
	LoginUC   *usecase.LoginUseCase
	LogoutUC  *usecase.LogoutUseCase
	RefreshUC *usecase.RefreshTokenUseCase
}

func NewAuthHandler(loginUC *usecase.LoginUseCase, logoutUC *usecase.LogoutUseCase, refreshUC *usecase.RefreshTokenUseCase) *AuthHandler {
	return &AuthHandler{LoginUC: loginUC, LogoutUC: logoutUC, RefreshUC: refreshUC}
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

	restutil.SetAuthCookie(w, result.Token, result.Expiration)

	rest.SendResponse(w, 200, "Login success")
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	identifier, ok := webutil.GetIdentifier(r)

	if !ok {
		rest.SendResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	h.LogoutUC.Execute(identifier)
	restutil.ClearAuthCookie(w)
	rest.SendResponse(w, 200, "Logout success")
}

func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	identifier, ok := webutil.GetIdentifier(r)

	if !ok {
		rest.SendResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	result := h.RefreshUC.Execute(identifier)
	if !result.Success {
		rest.SendResponse(w, http.StatusInternalServerError, result.Message)
		return
	}

	restutil.SetAuthCookie(w, result.Token, result.Expiration)

	rest.SendResponse(w, 200, "Refresh Token Success")
}
