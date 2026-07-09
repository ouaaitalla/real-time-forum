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
		helpers.Res(w, 400, "bad request", "error", logreq.Nickname, 0)
		return
	}
	// fmt.Println(logreq.Nickname)
	logreq.Nickname = strings.ToLower(logreq.Nickname)
	var hashedPassword string
	if (helpers.IsValidEmail(logreq.Nickname) && helpers.IsValidUserName(logreq.Nickname)){
		helpers.Res(w, 401, "Invalid Email Or Username", "error", "", 0)
		return
	}
	// if (helpers.IsValidPassword(logreq.Password)){
	// 	helpers.Res(w, 401, "Invalid Password Format", "error", "", 0)
	// 	return
	// }
	err = database.DB.QueryRow("SELECT id, password FROM users WHERE nickname = ? OR email = ?", logreq.Nickname, logreq.Nickname).Scan(&userID, &hashedPassword)
	if err != nil {
		fmt.Println(err)
		helpers.Res(w, 500, "user not found", "error", logreq.Nickname, 0)
		return
	}

	if !helpers.CheckPasswordHash(logreq.Password, hashedPassword) {
		helpers.Res(w, 401, "Invalid credentials", "error", logreq.Nickname, 0)
		return
	}
	sessionID, err := helpers.CreateSession(userID)
	if err != nil {
		helpers.Res(w, 500, "Failed to create session", "error", logreq.Nickname, 0)
		return
	}
	helpers.SetSessionCookie(w, sessionID)
	helpers.Res(w, 200, "complitly succes login", "success", logreq.Nickname, 0)
}
