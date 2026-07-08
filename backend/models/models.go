package models

type RegisterReq struct {
    Nickname string `json:"nickname"`
    Firstname string `json:"firstname"`
    Lastname string `json:"lastname"`
    Age string `json:"age"`
    Gender string `json:"gender"`
    Email string `json:"email"`
    Password string `json:"password"`
}

type LoginReq struct {
	// ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type Res struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Text     string `json:"text"`
}
type Post struct{
    ID  int `json:"id"`
    User_id int `json:"user_id"`
    Nickname string `json:"nickname"`
    Title string `json:"title"`
    Content string `json:"content"`
    Category string `json:"category"`
    Avatar string `json:"avatar"`
    Likes int `json:"likes"`
         int `json:"dislikes"`
    Created_at string `json:"created_at"`
}
type ReqCreatPost struct{
   Title string `json:"title"`
   Content string `json:"contend"`
   Category int `json:"category"`
}
type Reaction struct {
    ID           int    `json:"id"`
    UserID       int    `json:"user_id"`
    PostID       int    `json:"post_id"`
    ReactionType string `json:"reaction"`
}
type ReactionRequest struct {
    Reaction string `json:"reaction"`
}