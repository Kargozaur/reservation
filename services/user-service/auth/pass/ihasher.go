package pass

type IPwdHasher interface {
	Hash(password string) (string, error)
	VerifyPwd(password string, hash []byte) error
}
