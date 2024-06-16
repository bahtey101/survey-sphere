package service

import (
	"crypto/sha256"
	"fmt"
	"src/models"
	"src/package/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	salt       = "rand0m1237tr1ng"
	signingKey = "#new11k8y123423d#g43"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

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

func (service *AuthService) GenerateToken(_user models.User) (string, error) {
	user, err := service.repos.GetUser(models.User{
		Login:    _user.Login,
		Password: service.hashPassword(_user.Password)})
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		int(user.ID),
	})

	return token.SignedString([]byte(signingKey))
}

func (service *AuthService) hashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
