package service

type HashService interface {
	HashPassword(password string) (string, error)
	CompareHash(hash, password string) bool
}
