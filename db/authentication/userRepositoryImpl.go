package authentication

import (
	"core_chat/application/authentication/model"
	"core_chat/application/authentication/repository"
	"database/sql"
)

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) GetUserByUsername(username string) (model.User, error) {
	row := r.db.QueryRow("SELECT username, password, usertype FROM users WHERE username=?", username)
	var user model.User
	err := row.Scan(&user.UserName, &user.Password, &user.UserType)
	return user, err
}
