package authHandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"real-time-forum/backend/helpers"
	"real-time-forum/backend/models"
	"real-time-forum/database"
)
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var regReq models.RegisterReq
	err := json.NewDecoder(r.Body).Decode(&regReq)
	if err != nil {
		helpers.Res(w, 400, "Bad request try again", "error", regReq.Username)
		return
	}
	regReq.Email = strings.ToLower(regReq.Email)
	regReq.Username = strings.ToLower(regReq.Username)

	if !helpers.IsValidEmail(regReq.Email) {
		helpers.Res(w, 401, "Ivalid Email", "error", regReq.Username)
		return
	}

	if !helpers.IsValidUserName(regReq.Username) {
		helpers.Res(w, 401, "Ivalid Username", "error", regReq.Username)
		return
	}

	// if !helpers.IsValidPassword(regReq.Password) {
	// 	helpers.Res(w, 401, "Ivalid Password should be strong passoword", "error", regReq.Username)
	// 	return
	// }
	var existingEmail string
	err = database.DB.QueryRow("SELECT email FROM users WHERE email = ?", regReq.Email).Scan(&existingEmail)
	if err == nil {
		helpers.Res(w, 500, "User aleardy exist", "error", regReq.Username)
		return
	}
	hashedPassword, err := helpers.HashPassword(regReq.Password)
	if err != nil {
		helpers.Res(w, 500, "Internal server error", "error", regReq.Username)
		return
	}
	_, err = database.DB.Exec(
		"INSERT INTO users (age, nickname, email, password) VALUES (?, ?, ?, ?)",
		regReq.Age,
		regReq.Username,
		regReq.Email,
		hashedPassword,
	)
	if err != nil {
		fmt.Println(err)
		helpers.Res(w, 500, "Inserting User faild", "error", regReq.Username)
		return
	}
	helpers.Res(w, 200, "completely successful register", "success", regReq.Username)
}
