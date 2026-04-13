package request

type UpdatePasswordSchema struct {
	PasswordSchema
}

func (u *UpdatePasswordSchema) SwapPassword(password string) {
	u.Password = password
}
