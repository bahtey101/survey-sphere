package service

import (
	"src/models"
	"src/package/repository"
)

type AnswerService struct {
	repos     repository.Answers
	passRepos repository.Passes
}

func NewAnswerService(repos *repository.Answers, passRepos *repository.Passes) *AnswerService {
	return &AnswerService{repos: *repos, passRepos: *passRepos}
}

func (service *AnswerService) CreateAnswer(answer models.Answer) (*models.Answer, error) {
	return service.repos.CreateAnswer(answer)
}

func (service *AnswerService) GetAnswers(answer models.Answer) (*[]models.Answer, error) {
	return service.repos.GetAnswers(answer)
}

func (service *AnswerService) CreateAnswers(answers []models.Answer) (*[]models.Answer, error) {
	return service.repos.CreateAnswers(answers)
}
