package posthandlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"real-time-forum/backend/helpers"
	"real-time-forum/backend/models"
	"real-time-forum/database"
)

func CreatPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
	var req models.ReqCreatPost
	json.NewDecoder(r.Body).Decode(&req)
	userId, err := helpers.GetUserFromRequest(r)
	if err != nil {
		helpers.Res(w, 401, "Unauthorized", "error", "")
		return
	}
	var nickname string

	err = database.DB.QueryRow(
		"SELECT nickname FROM users WHERE id = ?",
		userId,
	).Scan(&nickname)
	if err != nil {
		helpers.Res(w, 500, "User not found", "error", "")
		return
	}
	_, err = database.DB.Exec(
		`INSERT INTO posts (user_id, nickname, title, content, category)
     VALUES (?, ?, ?, ?, ?)`,
		userId,
		nickname,
		req.Title,
		req.Content,
		req.Category,
	)
	if err != nil {
		fmt.Println(err)
		helpers.Res(w, 400, "Post created failed", "error", nickname)
		return
	}
	var post models.Post
	// post, err = GetPostHandler()
	fmt.Println("success")
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&post)
}
