package models

type SignUpRequest struct {
	UserName    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	PanCard     string `json:"panCard"`
}
