package service

import (
	"crypto/sha256"
	"fmt"
	"src/models"
	"src/package/repository"
)

const salt = "rand0m1237tr1ng"

type AuthService struct {
	repos repository.Authorization
}

func NewAuthService(repos *repository.Authorization) *AuthService {
	return &AuthService{repos: *repos}
}

func (service *AuthService) CreateUser(user models.User) (*models.User, error) {
	user.Password = service.hashPassword(user.Password)
	return service.repos.CreateUser(user)
}

func (service *AuthService) hashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
