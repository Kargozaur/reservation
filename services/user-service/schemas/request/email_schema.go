package request

type EmailSchema struct {
	Email string `json:"email" binding:"required,email"`
}
