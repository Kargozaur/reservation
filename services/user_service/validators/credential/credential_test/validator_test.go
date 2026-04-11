package credential_test

import (
	"errors"
	"testing"
	"user-service/validators/credential"
)

func TestValidateEmail(t *testing.T) {
	validator := credential.NewValidator()
	tests := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{
			name:    "valid email",
			email:   "test@example.com",
			wantErr: false,
		},
		{
			name:    "invalid email",
			email:   "invalid-email",
			wantErr: true,
		},
		{
			name:    "invalid email (no @)",
			email:   "example.com",
			wantErr: true,
		},
		{
			name:    "invalid email (no domain)",
			email:   "test@",
			wantErr: true,
		},
		{
			name:    "invalid character",
			email:   "test@@example.com",
			wantErr: true,
		},
		{
			name:    "empty email",
			email:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateEmail(tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateEmail(): expected %v, got %v", tt.wantErr, err)
			}
		})
	}
}

func TestValidator_ValidatePassword(t *testing.T) {
	v := credential.NewValidator()

	tests := []struct {
		name     string
		password string
		wantErrs int
	}{
		{
			name:     "matching",
			password: "Password123!",
			wantErrs: 0,
		},
		{
			name:     "short",
			password: "P1!",
			wantErrs: 1,
		},
		{
			name:     "no digit",
			password: "Password!",
			wantErrs: 1,
		},
		{
			name:     "no upper",
			password: "password123!",
			wantErrs: 1,
		},
		{
			name:     "no special",
			password: "Password123",
			wantErrs: 1,
		},
		{
			name:     "no special + no digit",
			password: "Password",
			wantErrs: 2,
		},
		{
			name:     "no upper + no special",
			password: "password123",
			wantErrs: 2,
		},
		{
			name:     "no upper + no digit",
			password: "password!",
			wantErrs: 2,
		},
		{
			name:     "empty",
			password: "",
			wantErrs: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := v.ValidatePassword(tt.password)
			if len(errs) != tt.wantErrs {
				t.Errorf("ValidatePassword(): got %d errors, want %d. Errors: %v", len(errs), tt.wantErrs, errors.Join(errs...))
			}
		})
	}
}
