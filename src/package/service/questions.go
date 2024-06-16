package service

import (
	"src/models"
	"src/package/repository"
)

type QuestionService struct {
	repos       repository.Questions
	surveyRepos repository.Surveys
}

func NewQuestionService(repos *repository.Questions, surveyRepos *repository.Surveys) *QuestionService {
	return &QuestionService{repos: *repos, surveyRepos: *surveyRepos}
}

func (service *QuestionService) CreateQuestion(question models.Question) (*models.Question, error) {
	return service.repos.CreateQuestion(question)
}

func (service *QuestionService) GetQuestion(question models.Question) (*models.Question, error) {
	return service.repos.GetQuestion(question)
}

func (service *QuestionService) GetQuestions(question models.Question) (*[]models.Question, error) {
	return service.repos.GetQuestions(question)
}

func (service *QuestionService) DeleteQuestion(question models.Question) (*models.Question, error) {
	return service.repos.DeleteQuestion(question)
}
