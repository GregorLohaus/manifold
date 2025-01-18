package types

type LoginData struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (v LoginData) GetMail() string {
	return v.Email
}

func (v LoginData) GetPass() string {
	return v.Password
}
