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
    Full_name string `json:"full_name"`
    Title string `json:"title"`
    Content string `json:"content"`
    Created_at string `json:"created_at"`
}
type ReqPost struct{
    
}