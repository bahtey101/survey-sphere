package repository

import (
	"src/models"

	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user models.User) (*models.User, error)
	GetUser(user models.User) (*models.User, error)

	// admin's method
	GetUsers() (*[]models.User, error)
	GetRole(user models.User) (*models.UserRole, error)
	DeleteUser(user models.User) (*models.User, error)
}

type Surveys interface {
	CreateSurvey(survey models.Survey) (*models.Survey, error)
	GetSurvey(survey models.Survey) (*models.Survey, error)
	GetSurveys(survey models.Survey) (*[]models.Survey, error)
	DeleteSurvey(survey models.Survey) (*models.Survey, error)
	GetSurveyAnswers(survey models.Survey) (*[]models.Answer, error)

	// admin's method
	GetAllSurveys() (*[]models.Survey, error)
}

type Questions interface {
	CreateQuestion(question models.Question) (*models.Question, error)
	GetQuestion(question models.Question) (*models.Question, error)
	GetQuestions(question models.Question) (*[]models.Question, error)
	DeleteQuestion(question models.Question) (*models.Question, error)
	CreateQuestions(questions []models.Question) (*[]models.Question, error)
}

type Passes interface {
	CreatePass(pass models.Pass) (*models.Pass, error)
	GetPass(pass models.Pass) (*models.Pass, error)
	GetPasses(pass models.Pass) (*[]models.Pass, error)
}

type Answers interface {
	CreateAnswer(answer models.Answer) (*models.Answer, error)
	GetAnswers(answer models.Answer) (*[]models.Answer, error)
	CreateAnswers([]models.Answer) (*[]models.Answer, error)
}

type Repository struct {
	Authorization
	Surveys
	Questions
	Passes
	Answers
}

func NewRepository(DB *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(DB),
		Surveys:       NewSurveyPostgres(DB),
		Questions:     NewQuestionPostgres(DB),
		Passes:        NewPassPostrgres(DB),
		Answers:       NewAnwserPosgtres(DB),
	}
}
