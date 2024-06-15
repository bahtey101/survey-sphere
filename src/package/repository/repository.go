package repository

import "gorm.io/gorm"

type Authorization interface{}

type Surveys interface{}

type Repository struct {
	Authorization
	Surveys
}

func NewRepository(DB *gorm.DB) *Repository {
	return &Repository{}
}
