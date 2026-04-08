package request

type PasswordSchema struct {
	Password string `json:"password" binding:"required"`
}
