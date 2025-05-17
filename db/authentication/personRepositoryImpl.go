package authentication

import (
	"core_chat/application/authentication/model"
	"core_chat/application/authentication/repository"
	"database/sql"
)

type personRepositoryImpl struct {
	db *sql.DB
}

func NewPersonRepository(db *sql.DB) repository.PersonRepository {
	return &personRepositoryImpl{db: db}
}

func (r *personRepositoryImpl) GetPersonByIdentifier(identifier string) (model.Person, error) {
	row := r.db.QueryRow("SELECT identifier, password, role FROM person WHERE identifier=?", identifier)
	var person model.Person
	err := row.Scan(&person.Identifier, &person.Password, &person.Role)
	return person, err
}
