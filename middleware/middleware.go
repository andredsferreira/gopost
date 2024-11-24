package middleware

import (
	"goweb01/service"
	"net/http"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "missing jwt cookie", http.StatusUnauthorized)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = service.VerifyJWT(cookie.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
