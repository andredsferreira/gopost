package main

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var HMACSecretKey []byte = []byte("f20jh6nbp3Vr0n2c02c0n1894j2vnrv0un2m40395jbv4j8v1pc2hu0489üvmcü319234")

func hashPassword(p string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), 12)
	if err != nil {
		log.Fatal(err)
	}
	return string(b), nil
}

func generateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"exp":      time.Now().Add(1 * time.Hour).Unix(),
			"username": username,
		})
	tokenString, err := token.SignedString(HMACSecretKey)
	if err != nil {
		log.Fatal(err)
	}
	return tokenString, nil
}
