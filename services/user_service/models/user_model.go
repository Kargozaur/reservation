package models

type User struct {
	BaseModel
	Email     string `gorm:"unique"`
	Password  string
	FirstName string
	LastName  string
	Token     []RefreshToken `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
