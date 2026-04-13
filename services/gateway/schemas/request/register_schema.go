package request

type RegisterSchema struct {
	EmailSchema
	PasswordSchema
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

func (r *RegisterSchema) SwapPassword(password string) {
	r.Password = password
}
