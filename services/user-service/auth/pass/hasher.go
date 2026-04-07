package pass

import "golang.org/x/crypto/bcrypt"

type Hasher struct {
	cost int
}

func NewHasher(cost int) *Hasher {
	return &Hasher{cost: cost}
}

func (h *Hasher) Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), h.cost)
}

func (h *Hasher) Verify(password string, hash []byte) error {
	return bcrypt.CompareHashAndPassword(hash, []byte(password))
}
