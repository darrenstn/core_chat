package service

type TokenService interface {
	GenerateToken(username string, userType int) (string, error)
	ValidateToken(tokenString string, requiredType int) (bool, string, int)
}
