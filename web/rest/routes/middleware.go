package routes

import (
	"core_chat/infra/serviceimpl"
	"core_chat/web/util"
	"net/http"
)

func Authenticate(next http.HandlerFunc, role string, requireEmailValidated bool) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}
		valid, identifier, _, isEmailValidated := serviceimpl.NewJWTTokenService().ValidateToken(cookie.Value, role)
		if !valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if requireEmailValidated != isEmailValidated {
			if !isEmailValidated {
				http.Error(w, "Email is not validated", http.StatusUnauthorized)
				return
			}
		}
		ctx := util.WithIdentifier(r.Context(), identifier)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
