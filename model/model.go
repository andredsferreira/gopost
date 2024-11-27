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

type Post struct {
	Username  string
	Title     string
	Content   string
	CreatedAt string
	Category  string
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

func GetAllPosts() ([]Post, error) {
	var posts []Post
	query := `
        SELECT username, title, content, created_at
        FROM posts
    `
	rows, err := db.MySql.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.Username, &p.Title, &p.Content, &p.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		posts = append(posts, p)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return posts, nil
}

func GetUserPosts(username string) ([]Post, error) {
	var posts []Post
	query := `
		SELECT username, title, content, created_at
		FROM posts
		WHERE username = ?;
	`
	rows, err := db.MySql.Query(query, username)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.Username, &p.Title, &p.Content, &p.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		posts = append(posts, p)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return posts, nil
}

func CreatePost(username, title, content string, categories []string) error {
	query := `
		INSERT INTO posts (username, title, content)
		VALUES (?, ?, ?)
	`
	result, err := db.MySql.Exec(query, username, title, content)
	if err != nil {
		return fmt.Errorf("failed to insert post: %w", err)
	}
	postID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID: %w", err)
	}
	for _, category := range categories {
		var categoryID int
		query := `
			SELECT id FROM categories WHERE category_name = ?
		`
		err := db.MySql.QueryRow(query, category).Scan(&categoryID)
		if err != nil {
			return fmt.Errorf("failed to find category %s: %w", category, err)
		}
		insertQuery := `
			INSERT INTO post_categories (post_id, category_id)
			VALUES (?, ?)
		`
		_, err = db.MySql.Exec(insertQuery, postID, categoryID)
		if err != nil {
			return fmt.Errorf("failed to insert post-category relation: %w", err)
		}
	}
	return nil
}

func GetAllCategories() ([]string, error) {
	var categories []string
	query := `
        SELECT category_name
        FROM categories
    `
	rows, err := db.MySql.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var category string
		if err := rows.Scan(&category); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return categories, nil
}
