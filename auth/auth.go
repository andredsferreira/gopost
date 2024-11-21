package auth

import (
	"errors"
	"fmt"
	"goweb01/data"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var HMACSecretKey []byte = []byte("f20jh6nbp3Vr0n2c02c0n1894j2vnrv0un2m40395jbv4j8v1pc2hu0489üvmcü319234")

func HashPassword(p string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), 12)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %v", err)
	}
	return string(b), nil
}

func CheckPasswordHash(p, h string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
	return err == nil
}

func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"expire":   time.Now().Add(1 * time.Hour).Unix(),
			"username": username,
		})
	tokenString, err := token.SignedString(HMACSecretKey)
	if err != nil {
		return "", fmt.Errorf("error signing token string: %v", err)
	}
	return tokenString, nil
}

func GetClaimsFromJWT(jwts string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(jwts, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v",
				token.Header["alg"])
		}
		return HMACSecretKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing jwt: %v", err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid jwt")
}

func Authorize(r *http.Request) error {
	username := r.FormValue("username")
	user, err := data.GetUserByUsername(username)
	if err != nil {
		return err
	}
	st, err := r.Cookie("session_token")
	if err != nil || st.Value != user.SessionToken || st.Value == "" {
		return fmt.Errorf("no session token provided: %v", err)
	}
	csrf := r.Header.Get("X-CSRF-Token")
	if csrf != user.CSRFToken || csrf == "" {
		return errors.New("csrf token header invalid")
	}
	return nil
}
