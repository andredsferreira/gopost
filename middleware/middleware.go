package middleware

import (
	"goweb01/service"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

func CreateStack(m ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(m) - 1; i >= 0; i-- {
			c := m[i]
			next = c(next)
		}
		return next
	}
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := time.Now()
		next.ServeHTTP(w, r)
		d := time.Since(s)
		log.Printf("INFO: %s %s %s\n", r.Method, r.URL.String(), d.String())
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
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
