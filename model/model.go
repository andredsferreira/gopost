package model

import (
	"database/sql"
	"fmt"
	"gopost/db"
	"regexp"
)

type User struct {
	Username string
	Password string
	Email    string
}

type Post struct {
	Username   string
	Title      string
	Content    string
	CreatedAt  string
	Categories []string
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
        SELECT *
        FROM posts
    `
	rows, err := db.MySql.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var p Post
		var id int
		if err := rows.Scan(&id, &p.Username, &p.Title, &p.Content, &p.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		postCategories, err := getPostCategories(id)
		if err != nil {
			return nil, fmt.Errorf("failed to execute query: %w", err)
		}
		p.Categories = postCategories
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
		SELECT *
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
		var id int
		if err := rows.Scan(&id, &p.Username, &p.Title, &p.Content, &p.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		postCategories, err := getPostCategories(id)
		if err != nil {
			return nil, fmt.Errorf("failed to execute query: %w", err)
		}
		p.Categories = postCategories
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

func getPostCategories(id int) ([]string, error) {
	var categories []string
	query := `
        SELECT c.category_name
        FROM categories c
        JOIN post_categories pc ON c.id = pc.category_id
        WHERE pc.post_id = ?
	`
	rows, err := db.MySql.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("error querying post categories: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var category string
		if err := rows.Scan(&category); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %v", err)
	}
	return categories, nil
}
