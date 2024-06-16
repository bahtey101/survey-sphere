package repository

import (
	"src/models"

	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user models.User) (*models.User, error)
	GetUser(user models.User) (*models.User, error)
}

type Surveys interface {
	CreateSurvey(userID int, survey models.Survey) (*models.Survey, error)
	GetSurvey(survey models.Survey) (*models.Survey, error)
	GetSurveys(userID int) (*[]models.Survey, error)
	DeleteSurvey(survey models.Survey) (*models.Survey, error)
}

type Questions interface {
	CreateQuestion(question models.Question) (*models.Question, error)
	GetQuestion(question models.Question) (*models.Question, error)
	GetQuestions(question models.Question) (*[]models.Question, error)
	DeleteQuestion(question models.Question) (*models.Question, error)
}

type Repository struct {
	Authorization
	Surveys
	Questions
}

func NewRepository(DB *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(DB),
		Surveys:       NewSurveyPostgres(DB),
	}
}
