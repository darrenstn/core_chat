package serviceimpl

import (
	"core_chat/application/authentication/service"
	"net/http"
	"time"
)

type HTTPServiceImpl struct{}

func NewHTTPService() service.HTTPService {
	return &HTTPServiceImpl{}
}

func (s *HTTPServiceImpl) SetAuthCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(5 * time.Minute),
		HttpOnly: true,
	})
}

func (s *HTTPServiceImpl) ClearAuthCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
	})
}
