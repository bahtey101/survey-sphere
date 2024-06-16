package repository

import (
	"src/models"

	"gorm.io/gorm"
)

type AuthPostgres struct {
	DB *gorm.DB
}

func NewAuthPostgres(DB *gorm.DB) *AuthPostgres {
	return &AuthPostgres{DB: DB}
}

func (postgres *AuthPostgres) CreateUser(user models.User) (*models.User, error) {
	err := postgres.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
