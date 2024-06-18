package service

import (
	"crypto/sha256"
	"errors"
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

func (service *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (service *AuthService) hashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (service *AuthService) GetUsers() (*[]models.User, error) {
	return service.repos.GetUsers()
}

func (service *AuthService) CheckRole(user models.User) bool {
	role, err := service.repos.GetRole(user)
	if err == nil && *role == models.Admin {
		return true
	}
	return false
}

func (service *AuthService) GetUser(user models.User) (*models.User, error) {
	return service.repos.GetUser(user)
}

func (service *AuthService) DeleteUser(user models.User) (*models.User, error) {
	return service.repos.DeleteUser(user)
}
