package models

type RegisterReq struct {
	Username string `json:"username"`
	Age      string `json:"age"`
	Email    string `json:"email"`
	Password string `json:"passowrd"`
}
type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"passowrd"`
}
type User struct {
	Username string `json:"username"`
	Age      string `json:"age"`
}
