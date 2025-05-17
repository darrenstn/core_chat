package routes

import (
	"core_chat/infra/serviceimpl"
	"core_chat/web/rest"
	"net/http"
)

func Authenticate(next http.HandlerFunc, role string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			rest.SendResponse(w, 400, "Missing token")
			return
		}
		valid, _, _ := serviceimpl.NewJWTTokenService().ValidateToken(cookie.Value, role)
		if !valid {
			rest.SendResponse(w, 401, "Unauthorized")
			return
		}
		next.ServeHTTP(w, r)
	})
}
