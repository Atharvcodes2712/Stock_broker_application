package models

type SignInRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
