package user

import (
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxentities"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxrepositories"
	"gorm.io/gorm"
)

type userRepository struct {
	dbClient *gorm.DB
}

func NewRepository(database *gorm.DB) Repository {
	return &userRepository{dbClient: database}
}

func (u userRepository) Insert(user entities.User) error {
	return u.dbClient.Create(&user).Error
}

func (u userRepository) Update(user *entities.User) error {
	return u.dbClient.Model(&user).Where(
		entities.User{Username: user.Username}).Updates(&user).Error
}

func (u userRepository) FindByUsername(username string) (*entities.User, error) {
	var user entities.User
	if err := u.dbClient.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, repositories.HandleError(err)
	}
	return &user, nil
}
