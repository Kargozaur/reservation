package models

import "github.com/google/uuid"

type RefreshToken struct {
	BaseModel
	Token  string
	UserID uuid.UUID
	User   User `gorm:"foreignKey:UserID"`
}
