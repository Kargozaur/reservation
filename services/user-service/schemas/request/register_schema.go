package request

import "user-service/models"

type RegisterSchema struct {
	EmailSchema
	PasswordSchema
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (r *RegisterSchema) ToUserModel() *models.User {
	return &models.User{
		Email:     r.Email,
		Password:  r.Password,
		FirstName: r.FirstName,
		LastName:  r.LastName,
	}
}

func (r *RegisterSchema) SwapPassword(password string) {
	r.Password = password
}
