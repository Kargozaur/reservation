package schemas

import "github.com/google/uuid"

type UserResponse struct {
	ID        uuid.UUID
	Email     string
	FirstName string
	LastName  string
}
