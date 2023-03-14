package model

type RegisterInfo struct {
	Username string `form:"username" json:"user"`
	Password string `form:"password" json:"password"`
	Email    string `form:"email" json:"email"`
}

type LoginInfo struct {
	Username string `form:"username" json:"user"`
	Password string `form:"password" json:"password"`
}

type EmailInfo struct {
	Email string
}
