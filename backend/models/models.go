package models

type Post struct {
	ID             string `json:"id"`
	UserID         string `json:"user_id"`
	AuthorNickname string `json:"author_nickname"`
	Title          string `json:"title"`
	Content        string `json:"content"`
	Category       string `json:"category"`
	CreatedAt      string `json:"created_at"`
	CommentCount   int    `json:"comment_count"`
}

type PostRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
}

