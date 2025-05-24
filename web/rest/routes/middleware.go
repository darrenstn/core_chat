package routes

import (
	"core_chat/infra/serviceimpl"
	"core_chat/web/rest"
	"core_chat/web/rest/util"
	"net/http"
)

func Authenticate(next http.HandlerFunc, role string, requireEmailValidated bool) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			rest.SendResponse(w, 400, "Missing token")
			return
		}
		valid, identifier, _, isEmailValidated := serviceimpl.NewJWTTokenService().ValidateToken(cookie.Value, role)
		if !valid {
			rest.SendResponse(w, 401, "Unauthorized")
			return
		}

		if requireEmailValidated != isEmailValidated {
			if !isEmailValidated {
				rest.SendResponse(w, 401, "Email is not validated")
				return
			}
		}
		ctx := util.WithIdentifier(r.Context(), identifier)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
