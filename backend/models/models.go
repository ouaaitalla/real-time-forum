package models

type RegisterReq struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
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
