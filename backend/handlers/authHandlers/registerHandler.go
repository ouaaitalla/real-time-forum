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
		fmt.Println(err)
		helpers.Res(w, 400, "Bad request try again", "error", regReq.Nickname, 0)
		return
	}
	regReq.Email = strings.ToLower(regReq.Email)
	regReq.Nickname = strings.ToLower(regReq.Nickname)
	if !helpers.IsValidEmail(regReq.Email) {
		helpers.Res(w, 401, "Ivalid Email", "error", regReq.Nickname, 0)
		return
	}

	if !helpers.IsValidUserName(regReq.Nickname) {
		helpers.Res(w, 401, "Ivalid Username", "error", regReq.Nickname, 0)
		return
	}
	if !helpers.IsValidPassword(regReq.Password) {
		helpers.Res(w, 401, "Ivalid Password should be strong passoword", "error", regReq.Nickname, 0)
		return
	}
	var existingEmail string
	err = database.DB.QueryRow("SELECT email FROM users WHERE email = ?", regReq.Email).Scan(&existingEmail)
	if err == nil {
		helpers.Res(w, 400, "User aleardy exist", "error", regReq.Nickname, 0)
		return
	}
	hashedPassword, err := helpers.HashPassword(regReq.Password)
	if err != nil {
		helpers.Res(w, 500, "Internal server error", "error", regReq.Nickname, 0)
		return
	}
	_, err = database.DB.Exec(
		"INSERT INTO users (age, nickname, email, password, firstname, lastname, gender) VALUES (?, ?, ?, ?, ?, ?, ?)",
		regReq.Age,
		regReq.Nickname,
		regReq.Email,
		hashedPassword,
		regReq.Firstname,
		regReq.Lastname,
		regReq.Gender,
	)
	if err != nil {
		fmt.Println(err)
		helpers.Res(w, 500, "Inserting User faild", "error", regReq.Nickname, 0)
		return
	}
	helpers.Res(w, 200, "completely successful register", "success", regReq.Nickname, 0)
}
