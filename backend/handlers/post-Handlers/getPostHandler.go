package posthandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"real-time-forum/backend/helpers"
	"real-time-forum/backend/models"
	"real-time-forum/database"
)

var Queryselector = "SELECT * FROM posts ORDER BY created_at DESC LIMIT ? OFFSET ?"

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	var limit int
	query := r.URL.Query()
	limitStr := query.Get("limit")
	cursorStr := query.Get("cursor")
	cur, err := strconv.Atoi(cursorStr)
	if err != nil {
		cur = 0
	}
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		limit = 20
	}
	rows, err1 := database.DB.Query(Queryselector, limit, cur)
	if err1 != nil {
		fmt.Println(err1)
		helpers.Res(w, 500, "Database Error", "error", "nope")
		return
	}
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.User_id, &post.Nickname, &post.Title, &post.Content, &post.Avatar, &post.Created_at)
		if err != nil {
			helpers.Res(w, 500, "Scan Error", "error", "nope")
			return
		}
		posts = append(posts, post)
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&posts)
}
