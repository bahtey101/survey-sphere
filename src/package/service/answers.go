package service

import (
	"src/models"
	"src/package/repository"
)

type AnswerService struct {
	repos repository.Answers
}

func NewAnswerService(repos *repository.Answers) *AnswerService {
	return &AnswerService{repos: *repos}
}

func (service *AnswerService) CreateAnswer(answer models.Answer) (*models.Answer, error) {
	return service.repos.CreateAnswer(answer)
}

func (service *AnswerService) GetAnswers(answer models.Answer) (*[]models.Answer, error) {
	return service.repos.GetAnswers(answer)
}
