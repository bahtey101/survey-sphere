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

func (postgres *AuthPostgres) GetUser(user models.User) (*models.User, error) {
	err := postgres.DB.Where(&user).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (postgres *AuthPostgres) GetUsers() (*[]models.User, error) {
	var users *[]models.User
	err := postgres.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (postgres *AuthPostgres) GetRole(user models.User) (*models.UserRole, error) {
	err := postgres.DB.First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user.Role, nil
}

func (postgres *AuthPostgres) DeleteUser(user models.User) (*models.User, error) {
	err := postgres.DB.Delete(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
