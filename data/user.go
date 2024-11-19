package data

import (
	"database/sql"
	"fmt"
	"goweb01/db"
)

type User struct {
	Username       string
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

func AddUser(username string, hashedPassword string) error {
	query := `
        INSERT INTO users (username, hashed_password) 
        VALUES (?, ?)
    `
	_, err := db.MySql.Exec(query, username, hashedPassword)
	if err != nil {
		return fmt.Errorf("AddUser: %v", err)
	}
	return nil
}

func GetUserByUsername(username string) (User, error) {
	var u User
	query := `
		SELECT *
		FROM users
		WHERE username = ?
	`
	row := db.MySql.QueryRow(query, username)
	if err := row.Scan(&u.Username, &u.HashedPassword,
		&u.SessionToken, &u.CSRFToken); err != nil {
		if err == sql.ErrNoRows {
			return u, fmt.Errorf("user does not exist")
		}
		return u, fmt.Errorf("getByUsername: %v", err)
	}
	return u, nil
}

func GetAllUsers() ([]User, error) {
	var users []User
	query := `
		SELECT * 
		FROM users
	`
	rows, err := db.MySql.Query(query)
	if err != nil {
		return nil, fmt.Errorf("getAll: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var cUser User
		if err := rows.Scan(&cUser.Username, &cUser.HashedPassword,
			&cUser.SessionToken, &cUser.CSRFToken); err != nil {
			return nil, fmt.Errorf("getAll: %v", err)
		}
		users = append(users, cUser)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getAll: %v", err)
	}
	return users, nil
}

func (u *User) UpdateUserSession(sessionToken, csrfToken string) error {
	query := `
		UPDATE users 
		SET session_token = ?, csrf_token = ?
		WHERE username = ?
	`
	_, err := db.MySql.Exec(query, sessionToken, csrfToken, u.Username)
	if err != nil {
		return fmt.Errorf("UpdateUserSession: %v", err)
	}
	return nil
}

func (u *User) ClearUserSession() error {
	query := `
		UPDATE users
		SET session_token = '', csrf_token = ''
		WHERE username = ?
	`
	_, err := db.MySql.Exec(query, u.Username)
	if err != nil {
		return fmt.Errorf("ClearUserSession: %v", err)
	}
	return nil
}
