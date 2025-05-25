package repository

import "core_chat/application/person/entity"

type PersonRepository interface {
	SavePerson(person entity.Person) error
	ExistsByIdentifier(identifier string) bool
	ExistsByEmail(email string) bool
}
