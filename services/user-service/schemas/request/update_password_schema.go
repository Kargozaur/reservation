package request

import "user-service/models"

type UpdatePasswordSchema struct {
	PasswordSchema
}

func (u *UpdatePasswordSchema) ToUserModel(user *models.User) {
	user.Password = u.Password
}
