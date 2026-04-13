package schemas

import (
	"net/mail"
	"user-service/models"
)

type UpdateEmail struct {
	Email string
}

func (u *UpdateEmail) ToModel(model *models.User) error {
	if _, err := mail.ParseAddress(u.Email); err != nil {
		return err
	}
	model.Email = u.Email
	return nil
}
