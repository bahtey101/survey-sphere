package service

import (
	"src/models"
	"src/package/repository"
)

type Authorization interface {
	CreateUser(user models.User) (*models.User, error)
	GenerateToken(user models.User) (string, error)
	ParseToken(token string) (int, error)
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

type Service struct {
	Authorization
	Surveys
	Questions
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(&repos.Authorization),
		Surveys:       NewSurveyService(&repos.Surveys),
	}
}
