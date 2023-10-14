package user

import "github.com/youtrolledhahaha/XDTROLLEDAxzxx/entities"

type Repository interface {
	Insert(user entities.User) error
	Update(user *entities.User) error
	FindByUsername(username string) (*entities.User, error)
}
