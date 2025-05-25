package person

import (
	"core_chat/application/person/entity"
	"core_chat/application/person/repository"
	"database/sql"
)

type personRepositoryImpl struct {
	db *sql.DB
}

func NewPersonRepository(db *sql.DB) repository.PersonRepository {
	return &personRepositoryImpl{db: db}
}

func (r *personRepositoryImpl) SavePerson(person entity.Person) error {
	query := `
		INSERT INTO person (identifier, password, role, name, email, date_of_birth, description, picture_path)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(
		query,
		person.Identifier,
		person.Password,
		person.Role,
		person.Name,
		person.Email,
		person.DateOfBirth,
		person.Description,
		person.PicturePath,
	)

	return err
}

func (r *personRepositoryImpl) ExistsByEmail(email string) bool {
	var existingEmail string
	err := r.db.QueryRow("SELECT email FROM person WHERE email = ?", email).Scan(&existingEmail)

	if err != nil {
		if err == sql.ErrNoRows {
			return false // Email not found — doesn't exist
		}
		return false // For safety, treat errors as "not existing"
	}

	return true // Email exists
}

func (r *personRepositoryImpl) ExistsByIdentifier(identifier string) bool {
	var existingIdentifier string
	err := r.db.QueryRow("SELECT identifier FROM person WHERE identifier = ?", identifier).Scan(&existingIdentifier)

	if err != nil {
		if err == sql.ErrNoRows {
			return false // Identifier not found — doesn't exist
		}
		return false // For safety, treat errors as "not existing"
	}

	return true // Identifier exists
}
