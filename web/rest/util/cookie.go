package util

import (
	"net/http"
	"time"
)

func SetAuthCookie(w http.ResponseWriter, token string, expTime time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  expTime,
		HttpOnly: true,
	})
}

func ClearAuthCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
	})
}
