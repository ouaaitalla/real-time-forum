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
	Username string `json:"username"`
	Password string `json:"password"`
}

type Res struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Text     string `json:"text"`
}
