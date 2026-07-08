package db

import (
	"database/sql"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	_ "modernc.org/sqlite"

	"real-time-forum/backend/models"
)

var Conn *sql.DB

func Init(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	database, err := sql.Open("sqlite", path)
	if err != nil {
		return err
	}
	database.SetMaxOpenConns(1)
	if _, err := database.Exec(`PRAGMA foreign_keys = ON;`); err != nil {
		return err
	}
	Conn = database
	return migrate()
}

func migrate() error {
	schema := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			nickname TEXT NOT NULL UNIQUE,
			age INTEGER NOT NULL,
			gender TEXT NOT NULL,
			first_name TEXT NOT NULL,
			last_name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password_hash TEXT NOT NULL,
			created_at TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS sessions (
			id TEXT PRIMARY KEY,
			user_id TEXT NOT NULL,
			created_at TEXT NOT NULL,
			expires_at TEXT NOT NULL,
			FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
		);`,
		`CREATE TABLE IF NOT EXISTS posts (
			id TEXT PRIMARY KEY,
			user_id TEXT NOT NULL,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			category TEXT NOT NULL,
			created_at TEXT NOT NULL,
			FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
		);`,
		`CREATE TABLE IF NOT EXISTS comments (
			id TEXT PRIMARY KEY,
			post_id TEXT NOT NULL,
			user_id TEXT NOT NULL,
			content TEXT NOT NULL,
			created_at TEXT NOT NULL,
			FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE,
			FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
		);`,
		`CREATE TABLE IF NOT EXISTS messages (
			id TEXT PRIMARY KEY,
			sender_id TEXT NOT NULL,
			receiver_id TEXT NOT NULL,
			content TEXT NOT NULL,
			created_at TEXT NOT NULL,
			FOREIGN KEY(sender_id) REFERENCES users(id) ON DELETE CASCADE,
			FOREIGN KEY(receiver_id) REFERENCES users(id) ON DELETE CASCADE
		);`,
		`CREATE INDEX IF NOT EXISTS idx_posts_created_at ON posts(created_at DESC);`,
		`CREATE INDEX IF NOT EXISTS idx_comments_post_created ON comments(post_id, created_at ASC);`,
		`CREATE INDEX IF NOT EXISTS idx_messages_pair_created ON messages(sender_id, receiver_id, created_at DESC);`,
		`CREATE INDEX IF NOT EXISTS idx_sessions_user ON sessions(user_id);`,
	}
	for _, stmt := range schema {
		if _, err := Conn.Exec(stmt); err != nil {
			return err
		}
	}
	return nil
}

func now() string {
	return time.Now().UTC().Format(time.RFC3339Nano)
}



func CreatePost(userID string, input models.PostRequest) (models.Post, error) {
	post := models.Post{ID: uuid.NewString(), UserID: userID, Title: strings.TrimSpace(input.Title), Content: strings.TrimSpace(input.Content), Category: strings.TrimSpace(input.Category), CreatedAt: now()}
	_, err := Conn.Exec(`INSERT INTO posts (id, user_id, title, content, category, created_at) VALUES (?, ?, ?, ?, ?, ?)`, post.ID, post.UserID, post.Title, post.Content, post.Category, post.CreatedAt)
	return post, err
}

func Posts() ([]models.Post, error) {
	rows, err := Conn.Query(`SELECT p.id, p.user_id, u.nickname, p.title, p.content, p.category, p.created_at,
		(SELECT COUNT(*) FROM comments c WHERE c.post_id = p.id) AS comment_count
		FROM posts p JOIN users u ON u.id = p.user_id ORDER BY p.created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.AuthorNickname, &post.Title, &post.Content, &post.Category, &post.CreatedAt, &post.CommentCount); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, rows.Err()
}
