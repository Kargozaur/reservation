package credential

import (
	"fmt"
	"net/mail"
	"unicode"
)

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}
func (v *Validator) ValidatePassword(password string) []error {
	var res []error
	var isSpecial, isUpper, isDigit bool
	if len(password) < 8 {
		res = append(res, fmt.Errorf("Password must be at least 8 characters"))
	}
	for _, ch := range password {
		switch {
		case unicode.IsDigit(ch):
			isDigit = true
		case unicode.IsUpper(ch):
			isUpper = true
		case v.isSpecial(ch):
			isSpecial = true
		}
	}
	if !isDigit {
		res = append(res, fmt.Errorf("Password must contain at least one digit"))
	}
	if !isUpper {
		res = append(res, fmt.Errorf("Password must contain at least one uppercase letter"))
	}
	if !isSpecial {
		res = append(res, fmt.Errorf("Password must contain at least one special character"))
	}
	return res
}

func (v *Validator) isSpecial(ch rune) bool {
	switch ch {
	case '!', '#', '$', '%', '&', '*', '+', '-', '.', '/', ':', ';', '<', '=', '>', '?', '@', '^', '_', '`', '|', '~':
		return true
	default:
		return false
	}
}
