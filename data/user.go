package data

import (
	"database/sql"
	"fmt"
	"goweb01/db"
)

type User struct {
	Username string
	Password string
}

func AddUser(username string, hashedPassword string) error {
	query := `
        INSERT INTO users (username, hashed_password) 
        VALUES (?, ?)
    `
	_, err := db.MySql.Exec(query, username, hashedPassword)
	if err != nil {
		return fmt.Errorf("addUser: %v", err)
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
	if err := row.Scan(&u.Username, &u.Password); err != nil {
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
		if err := rows.Scan(&cUser.Username, &cUser.Password); err != nil {
			return nil, fmt.Errorf("getAll: %v", err)
		}
		users = append(users, cUser)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getAll: %v", err)
	}
	return users, nil
}
