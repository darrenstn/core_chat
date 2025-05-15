package repository

import "core_chat/application/authentication/model"

type UserRepository interface {
	GetUserByUsername(username string) (model.User, error)
}
