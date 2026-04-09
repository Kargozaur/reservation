package pass_test

import (
	"testing"
	"user-service/auth/pass"
)

func TestHasher(t *testing.T) {
	hasher := pass.NewHasher(10)
	tests := []struct {
		name     string
		password string
	}{
		{name: "test1", password: "password1"},
		{name: "test2", password: "password2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := hasher.Hash(tt.password)
			if err != nil {
				t.Fatal(err)
			}
			if err := hasher.Verify(tt.password, []byte(hash)); err != nil {
				t.Fatal(err)
			}
		})
	}
}
