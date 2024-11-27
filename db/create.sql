CREATE TABLE users (
    username VARCHAR(255) PRIMARY KEY,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE
);
CREATE TABLE posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255),
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (username) REFERENCES users(username) ON DELETE CASCADE
);
CREATE TABLE categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    category_name VARCHAR(60) NOT NULL UNIQUE
);
CREATE TABLE post_categories (
    post_id INT,
    category_id INT,
    PRIMARY KEY (post_id, category_id),
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);
INSERT INTO categories (category_name)
VALUES ('Programming'),
    ('Literature'),
    ('Technology'),
    ('Science'),
    ('Health'),
    ('Education'),
    ('Travel'),
    ('Food'),
    ('Art'),
    ('Music'),
    ('Sports'),
    ('Finance'),
    ('Business'),
    ('Lifestyle'),
    ('Photography'),
    ('Movies'),
    ('History'),
    ('Gaming'),
    ('Environment'),
    ('Politics');
