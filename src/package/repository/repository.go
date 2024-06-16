package repository

import (
	"src/models"

	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user models.User) (*models.User, error)
	GetUser(user models.User) (*models.User, error)
}

type Surveys interface{}

type Repository struct {
	Authorization
	Surveys
}

func NewRepository(DB *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(DB),
	}
}
