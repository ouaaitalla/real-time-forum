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
    Created_at string `json:"created_at"`
}
type ReqCreatPost struct{
   Title string `json:"title"`
   Content string `json:"contend"`
   Category int `json:"category"`
}