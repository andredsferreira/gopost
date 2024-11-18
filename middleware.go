package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func verifyJWT(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] == nil {
			log.Fatal("no token in the request header")
		}
		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return HMACSecretKey, nil
		})
		if err != nil {
			log.Fatal("error parsing token", err)
		}
		_, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Fatal("token claims are not valid ")
		}
		if !token.Valid {
			log.Fatal("token is not valid")
		}
		next.ServeHTTP(w, r)
	})
}
