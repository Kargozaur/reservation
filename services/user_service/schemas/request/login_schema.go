package request

type LoginSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
