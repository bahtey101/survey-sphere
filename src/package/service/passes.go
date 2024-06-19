package service

import (
	"src/models"
	"src/package/repository"
)

type PassService struct {
	repos repository.Passes
}

func NewPassService(repos *repository.Passes) *PassService {
	return &PassService{repos: *repos}
}

func (service *PassService) CreatePass(pass models.Pass) (*models.Pass, error) {
	return service.repos.CreatePass(pass)
}

func (service *PassService) GetPass(pass models.Pass) (*models.Pass, error) {
	return service.repos.GetPass(pass)
}

func (service *PassService) GetPasses(pass models.Pass) (*[]models.Pass, error) {
	return service.repos.GetPasses(pass)
}

func (service *PassService) GetUserPassCount(user models.User) int {
	return service.repos.GetUserPassCount(user)
}
