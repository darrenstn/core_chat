package service

type ValidatorService interface {
	IsIdentifierValid(identifier string) bool
	IsEmailValid(email string) bool
	IsPasswordValid(password string) bool
}
