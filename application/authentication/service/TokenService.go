package service

type TokenService interface {
	GenerateToken(identifier string, role string) (string, error)
	ValidateToken(tokenString string, role string) (bool, string, string)
}
