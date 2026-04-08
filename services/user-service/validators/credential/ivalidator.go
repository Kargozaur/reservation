package credential

type IValidator interface {
	ValidateEmail(email string) error
	ValidatePassword(password string) []error
}
