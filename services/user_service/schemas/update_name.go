package schemas

import "user-service/models"

type UpdateName struct {
	FirstName *string
	LastName  *string
}

func (u *UpdateName) ToModel(model *models.User) {
	if u.FirstName != nil {
		model.FirstName = *u.FirstName
	}
	if u.LastName != nil {
		model.LastName = *u.LastName
	}
}
