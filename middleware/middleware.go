package middleware

import (
	"fmt"
	"gopost/service"
	"log"
	"net/http"
	"time"
)

type customResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (cw *customResponseWriter) WriteHeader(statusCode int) {
	cw.statusCode = statusCode
	cw.ResponseWriter.WriteHeader(statusCode)
}

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cw := &customResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		s := time.Now()
		next.ServeHTTP(cw, r)
		d := time.Since(s).Seconds()
		log.Printf("HTTP %s %s %d %fs\n", r.Method, r.URL.String(), cw.statusCode, d)
	}
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")
		if err != nil {
			if err == http.ErrNoCookie {
				fmt.Fprint(w, "<p>You must login first</p>")
				http.Error(w, "missing jwt cookie", http.StatusUnauthorized)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = service.VerifyJWT(cookie.Value)
		if err != nil {
			fmt.Fprint(w, "<p>You must login first</p>")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
