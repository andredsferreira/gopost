package model

import (
	"database/sql"
	"fmt"
	"goweb01/db"
	"regexp"
)

type User struct {
	Username string
	Password string
	Email    string
}

func GetUserByUsername(username string) (User, error) {
	var u User
	query := `
		SELECT username, password, email
		FROM users
		WHERE username = ?
	`
	row := db.MySql.QueryRow(query, username)
	if err := row.Scan(&u.Username, &u.Password, &u.Email); err != nil {
		if err == sql.ErrNoRows {
			return u, fmt.Errorf("user does not exist")
		}
		return u, err
	}
	return u, nil
}

func AddUser(username, password, email string) error {
	query := `
        INSERT INTO users (username, password, email) 
        VALUES (?, ?, ?)
    `
	_, err := db.MySql.Exec(query, username, password, email)
	if err != nil {
		return err
	}
	return nil
}

func ValidateUser(username, password, email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	isValidEmail := regexp.MustCompile(emailRegex).MatchString(email)
	if len(username) >= 3 && len(password) >= 3 && isValidEmail {
		return true
	}
	return false
}
