package repository

import (
	"core_chat/application/person/entity"
	"core_chat/application/person/model"
)

type PersonRepository interface {
	SavePerson(person entity.Person) error
	ExistsByIdentifier(identifier string) bool
	ExistsByEmail(email string) bool
	FindProfileByIdentifier(identifier string) (*model.Profile, error)
}
