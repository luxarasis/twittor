package middlew

import (
	"net/http"

	"github.com/luxarasis/twittor/jwt"
)

func ValidateJWT(nex http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := jwt.ProccesJWTToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "JWT Token invalid: "+err.Error(), http.StatusBadRequest)
			return
		}

		nex.ServeHTTP(w, r)
	}
}
