package request

import "user-service/models"

type UpdateNameSchema struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
}

func (s *UpdateNameSchema) ToUserModel(user *models.User) {
	if s.FirstName != nil {
		user.FirstName = *s.FirstName
	}
	if s.LastName != nil {
		user.LastName = *s.LastName
	}
}
