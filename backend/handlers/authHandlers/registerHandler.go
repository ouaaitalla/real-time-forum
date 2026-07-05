package authHandlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"real-time-forum/backend/helpers"
	"real-time-forum/backend/models"
	"real-time-forum/database"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("welcome in register")
	var regReq models.RegisterReq
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&regReq)
	if !helpers.IsValidEmail(regReq.Email) || !helpers.IsValidUserName(regReq.Username) || !helpers.IsValidPassword(regReq.Password) {
		//err
		return
	}
	var existingEmail string
	err = database.DB.QueryRow("SELECT email FROM users WHERE email = ?", regReq.Email).Scan(&existingEmail)
	if err == nil {
		//err
		return
	} else if err != sql.ErrNoRows {
		//err
		return
	}
	hashedPassword, err := helpers.HashPassword(regReq.Password)
	if err != nil {
		//err
		return
	}
	_, err = database.DB.Exec(
		"INSERT INTO users (age, nickname, email, password) VALUES (?, ?, ?, ?)",
		regReq.Age, regReq.Username, hashedPassword,
	)
	user.Username = regReq.Username
	user.Age = regReq.Age
	err = json.NewEncoder(w).Encode(user)
}
