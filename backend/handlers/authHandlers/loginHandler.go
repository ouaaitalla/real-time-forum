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
	if err != nil{
		helpers.Res(w, 400, "bad request", "error", logreq.Username)
		return
	}
	fmt.Println(logreq.Username)
	logreq.Username = strings.ToLower(logreq.Username)
	var hashedPassword string
	err = database.DB.QueryRow("SELECT id, password FROM users WHERE nickname = ?", logreq.Username).Scan(&userID, &hashedPassword)
	if err != nil {
		fmt.Println(err)
		helpers.Res(w, 500, "user not found", "error", logreq.Username)
		return
	}

	if !helpers.CheckPasswordHash(logreq.Password, hashedPassword) {
		helpers.Res(w, 401, "Invalid credentials", "error", logreq.Username)
		return
	}
	sessionID, err := helpers.CreateSession(userID)
	if err != nil {
		helpers.Res(w, 500, "Failed to create session", "error", logreq.Username)
		return
	}
	helpers.SetSessionCookie(w, sessionID)
	helpers.Res(w, 200, "complitly succes login", "success", logreq.Username)
}
