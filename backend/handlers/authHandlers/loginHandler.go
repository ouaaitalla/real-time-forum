package authHandlers

import (
	// "database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"real-time-forum/backend/helpers"
	"real-time-forum/backend/models"
	"real-time-forum/database"
)

var userID int

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var logreq models.LoginReq
	err := json.NewDecoder(r.Body).Decode(&logreq)
	if err != nil {
		fmt.Println(err)
		helpers.Res(w, 400, "bad request", "error", logreq.Nickname)
		return
	}
	fmt.Println(logreq.Nickname)
	logreq.Nickname = strings.ToLower(logreq.Nickname)
	var hashedPassword string
	err = database.DB.QueryRow("SELECT id, password FROM users WHERE nickname = ?", logreq.Nickname).Scan(&userID, &hashedPassword)
	if err != nil {
		fmt.Println(err)
		helpers.Res(w, 500, "user not found", "error", logreq.Nickname)
		return
	}

	if !helpers.CheckPasswordHash(logreq.Password, hashedPassword) {
		helpers.Res(w, 401, "Invalid credentials", "error", logreq.Nickname)
		return
	}
	sessionID, err := helpers.CreateSession(userID)
	if err != nil {
		helpers.Res(w, 500, "Failed to create session", "error", logreq.Nickname)
		return
	}
	helpers.SetSessionCookie(w, sessionID)
	helpers.Res(w, 200, "complitly succes login", "success", logreq.Nickname)
}
