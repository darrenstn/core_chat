package service

type TokenService interface {
	GenerateToken(identifier string, role string, emailValidated bool) (string, error)
	ValidateToken(tokenString string, role string, emailValidated bool) (bool, string, string, bool)
}
