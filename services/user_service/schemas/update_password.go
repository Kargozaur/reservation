package schemas

import "user-service/models"

type UpdatePassword struct {
	Password string
}

func (u *UpdatePassword) SwapPassword(hashedPassword string) *UpdatePassword {
	u.Password = hashedPassword
	return u
}

func (u *UpdatePassword) ToModel(user *models.User) {
	user.Password = u.Password
}
