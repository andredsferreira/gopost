package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var HMACSecretKey []byte = []byte("c058135f3d1ddfba6c62026751d04ae907200201ca4f24c4e908f5b103d1abc9")

func HashPassword(p string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), 12)
	if err != nil {
		return "", fmt.Errorf("generating hash from password: %w", err)
	}
	return string(b), nil
}

func CheckPasswordHash(p, h string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
	return err == nil
}

func GenerateJWT(u string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": u,
			"expire":   time.Now().Add(10 * time.Minute).Unix(),
		})
	ts, err := t.SignedString(HMACSecretKey)
	if err != nil {
		return "", fmt.Errorf("signing token string: %w", err)
	}
	return ts, nil
}

func VerifyJWT(ts string) error {
	token, err := jwt.Parse(ts, func(token *jwt.Token) (interface{}, error) {
		return HMACSecretKey, nil
	})
	if err != nil {
		return fmt.Errorf("error parsing jwt: %w", err)
	}
	if !token.Valid {
		return fmt.Errorf("invalid jwt")
	}
	return nil
}

func GetClaimsFromJWT(ts string) (jwt.MapClaims, error) {
	t, err := jwt.Parse(ts, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return HMACSecretKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing jwt: %w", err)
	}
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid jwt")
}

func GetUsernameFromCookie(c *http.Cookie) string {
	claims, err := GetClaimsFromJWT(c.Value)
	if err != nil {
		return ""
	}
	return claims["username"].(string)
}
