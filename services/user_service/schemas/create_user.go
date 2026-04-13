package schemas

import "user-service/models"

type CreateUser struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
}

func (c CreateUser) ToModel() *models.User {
	return &models.User{
		Email:     c.Email,
		Password:  c.Password,
		FirstName: c.FirstName,
		LastName:  c.LastName,
	}
}

func (c *CreateUser) SwapPassword(hashedPassword string) *CreateUser {
	c.Password = hashedPassword
	return c
}
