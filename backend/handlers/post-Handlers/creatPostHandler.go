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
	errD := json.NewDecoder(r.Body).Decode(&req)
	if errD != nil{
		helpers.Res(w, 400, "bad request", "error", "", 0)
		return
	}
	userId, err := helpers.GetUserFromRequest(r)
	if err != nil {
		helpers.Res(w, 401, "Unauthorized", "error", "", 0)
		return
	}
	var nickname string

	err = database.DB.QueryRow(
		"SELECT nickname FROM users WHERE id = ?",
		userId,
	).Scan(&nickname)
	if err != nil {
		helpers.Res(w, 500, "User not found", "error", "", 0)
		return
	}
	result, err1 := database.DB.Exec(
		`INSERT INTO posts (user_id, nickname, title, content, category)
     VALUES (?, ?, ?, ?, ?)`,
		userId,
		nickname,
		req.Title,
		req.Content,
		req.Category,
	)
	if err1 != nil {
		fmt.Println(err)
		helpers.Res(w, 400, "Post created failed", "error", nickname, 0)
		return
	}
	var id int64
	id, err = result.LastInsertId()
	if err != nil{
		helpers.Res(w, 500, "error get id", "error", "", 0)
		return
	}
	helpers.Res(w, 200, "completly created post sussfuly", "success", "", id)
}
