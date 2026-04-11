package models

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *BaseModel) BeforeCreate() error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	b.ID = id
	now := time.Now().UTC()
	b.CreatedAt = now
	b.UpdatedAt = now
	return nil
}
func (b *BaseModel) BeforeUpdate() error {
	now := time.Now().UTC()
	b.UpdatedAt = now
	return nil
}
