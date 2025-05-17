package repository

import "core_chat/application/authentication/model"

type PersonRepository interface {
	GetPersonByIdentifier(Identifier string) (model.Person, error)
}
