package helpers

import (
	// "database/sql"
	"database/sql"
	"fmt"
	"net/http"
	"real-time-forum/database"
	"time"

	"github.com/google/uuid"
)

const SESSION_COOKIE_NAME = "forum_session"

func CreateSession(userID int) (string, error) {
	sessionID := uuid.New().String()
	expiresAt := time.Now().Add(24 * time.Hour)

	query := `INSERT INTO sessions (id, user_id, expires_at) VALUES (?, ?, ?)`
	_, err := database.DB.Exec(query, sessionID, userID, expiresAt)
	if err != nil {
		return "", fmt.Errorf("failed to create session: %w", err)
	}

	return sessionID, nil
}

func SetSessionCookie(w http.ResponseWriter, sessionID string) {
	http.SetCookie(w, &http.Cookie{
		Name:     SESSION_COOKIE_NAME,
		Value:    sessionID,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})
}

func GetUserIDFromSession(sessionID string) (int, error) {
	var userID int
	var expiresAt time.Time

	query := `SELECT user_id, expires_at FROM sessions WHERE id = ?`
	err := database.DB.QueryRow(query, sessionID).Scan(&userID, &expiresAt)
	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("session not found")
	}
	if err != nil {
		return 0, fmt.Errorf("database error: %w", err)
	}

	if time.Now().After(expiresAt) {
		DeleteSession(sessionID)
		return 0, fmt.Errorf("session expired")
	}

	return userID, nil
}

func DeleteSession(sessionID string) error {
	query := `DELETE FROM sessions WHERE id = ?`
	_, err := database.DB.Exec(query, sessionID)
	return err
}

func GetUserFromRequest(r *http.Request) (int, error) {
	cookie, err := r.Cookie(SESSION_COOKIE_NAME)
	if err != nil {
		return 0, err
	}
	return GetUserIDFromSession(cookie.Value)
}

func LoggedIn(r *http.Request) bool {
	id, err := GetUserFromRequest(r)
	return err == nil && id > 0
}

func ValidSession(r *http.Request) bool {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return false
	}

	var userID int
	var expiresAt time.Time

	err = database.DB.QueryRow(`
		SELECT user_id, expires_at
		FROM sessions
		WHERE id = ?
	`, cookie.Value).Scan(&userID, &expiresAt)
	if err != nil {
		return false
	}

	// valid only if NOT expired
	if time.Now().After(expiresAt) {
		return false
	}

	return true
}
