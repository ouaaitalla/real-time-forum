package models

type registerReq struct{
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"passowrd"`
}
type loginReq struct{
	Username string `json:"username"`
	Password string `json:"passowrd"`
}