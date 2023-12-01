package dto_user

type RequestLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ResponseLogin struct {
	Token string `json:"token"`
}
