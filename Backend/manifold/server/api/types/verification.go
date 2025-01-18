package types

type Verification struct {
	Email           string  `json:"email" validate:"required,email"`
	Password        string  `json:"password" validate:"required"`
	RegistrationKey *string `json:"registration_key,omitempty"`
}

func (v Verification) GetMail() string {
	return v.Email
}

func (v Verification) GetPass() string {
	return v.Password
}
