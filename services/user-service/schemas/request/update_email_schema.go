package request

import "user-service/models"

type UpdateEmailSchema struct {
	EmailSchema
}

func (u *UpdateEmailSchema) ToUserModel(user *models.User) {
	user.Email = u.Email
}
