package main

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func formatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func hashPassword(p string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	if err != nil {
		log.Fatal(err)
	}
	return string(b), nil
}
