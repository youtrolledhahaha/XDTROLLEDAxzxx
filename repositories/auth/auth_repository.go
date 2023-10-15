package auth

import (
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxentities"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxrepositories"
	"gorm.io/gorm"
)

type authRepository struct {
	dbClient *gorm.DB
}

func NewRepository(dbClient *gorm.DB) Repository {
	return &authRepository{dbClient: dbClient}
}

func (a authRepository) Insert(auth entities.Auth) error {
	return a.dbClient.Create(&auth).Error
}

func (a authRepository) Update(auth *entities.Auth) error {
	return a.dbClient.Model(&auth).Updates(&auth).Error
}

func (a authRepository) GetFirst() (*entities.Auth, error) {
	var auth entities.Auth
	tx := a.dbClient.Find(&auth)
	if tx.Error != nil {
		return nil, repositories.HandleError(tx.Error)
	}
	if tx.RowsAffected == 0 {
		return nil, repositories.ErrNotFound
	}
	return &auth, nil
}
