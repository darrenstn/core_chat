package dto

type Claims struct {
	Identifier     string
	Role           string
	EmailValidated bool
	ExpiresAt      int64
}
