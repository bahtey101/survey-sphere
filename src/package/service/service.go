package service

import (
	"src/models"
	"src/package/repository"
)

type Authorization interface {
	CreateUser(user models.User) (*models.User, error)
	GenerateToken(user models.User) (string, error)
	ParseToken(token string) (int, error)

	// admin's method
	GetUsers() (*[]models.User, error)
	CheckRole(user models.User) bool
	DeleteUser(user models.User) (*models.User, error)
}

type Surveys interface {
	CreateSurvey(survey models.Survey) (*models.Survey, error)
	GetSurvey(survey models.Survey) (*models.Survey, error)
	GetSurveys(survey models.Survey) (*[]models.Survey, error)
	DeleteSurvey(survey models.Survey) (*models.Survey, error)
	GetSurveyPasses(survey models.Survey) (*[]models.Pass, error)

	// admin's method
	GetAllSurveys() (*[]models.Survey, error)
}

type Questions interface {
	CreateQuestion(question models.Question) (*models.Question, error)
	GetQuestion(question models.Question) (*models.Question, error)
	GetQuestions(question models.Question) (*[]models.Question, error)
	DeleteQuestion(question models.Question) (*models.Question, error)
	CreateQuestions([]models.Question) (*[]models.Question, error)
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

type Service struct {
	Authorization
	Surveys
	Questions
	Passes
	Answers
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(&repos.Authorization),
		Surveys:       NewSurveyService(&repos.Surveys),
		Questions:     NewQuestionService(&repos.Questions, &repos.Surveys),
		Passes:        NewPassService(&repos.Passes),
		Answers:       NewAnswerService(&repos.Answers, &repos.Passes),
	}
}
