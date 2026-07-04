package posthandlers

import (
	"net/http"
	"strings"

	"real-time-forum/backend/db"
	"real-time-forum/backend/httpx"
	"real-time-forum/backend/middleware"
	"real-time-forum/backend/models"
)

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		posts, err := db.Posts()
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "could not load posts")
			return
		}
		httpx.JSON(w, http.StatusOK, map[string]any{"posts": posts})
	case http.MethodPost:
		user, ok := middleware.CurrentUser(r)
		if !ok {
			httpx.Error(w, http.StatusUnauthorized, "authentication required")
			return
		}
		var input models.PostRequest
		if err := httpx.Decode(r, &input); err != nil {
			httpx.Error(w, http.StatusBadRequest, "invalid JSON body")
			return
		}
		if msg := validatePost(input); msg != "" {
			httpx.Error(w, http.StatusBadRequest, msg)
			return
		}
		post, err := db.CreatePost(user.ID, input)
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "could not create post")
			return
		}
		post.AuthorNickname = user.Nickname
		httpx.JSON(w, http.StatusCreated, map[string]any{"post": post})
	default:
		httpx.Error(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func validatePost(input models.PostRequest) string {
	if len(strings.TrimSpace(input.Title)) < 3 {
		return "title must be at least 3 characters"
	}
	if len(strings.TrimSpace(input.Content)) < 5 {
		return "content must be at least 5 characters"
	}
	if strings.TrimSpace(input.Category) == "" {
		return "category is required"
	}
	return ""
}
