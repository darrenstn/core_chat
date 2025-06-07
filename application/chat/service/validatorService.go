package service

type ValidatorService interface {
	IsIdentifierValid(identifier string) bool
}
