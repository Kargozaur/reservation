package pass

type IPwdHasher interface {
	Hash(password string) (string, error)
	Verify(password string, hash []byte) error
}
