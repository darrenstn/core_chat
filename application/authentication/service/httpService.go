package service

import "net/http"

type HTTPService interface {
	SetAuthCookie(w http.ResponseWriter, token string)
	ClearAuthCookie(w http.ResponseWriter)
}
