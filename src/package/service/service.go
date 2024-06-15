package service

import "src/package/repository"

type Authorization interface{}

type Surveys interface{}

type Service struct {
	Authorization
	Surveys
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
