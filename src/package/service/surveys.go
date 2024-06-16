package service

import (
	"src/models"
	"src/package/repository"
)

type SurveyService struct {
	repos repository.Surveys
}

func NewSurveyService(repos *repository.Surveys) *SurveyService {
	return &SurveyService{repos: *repos}
}

func (service *SurveyService) CreateSurvey(userID int, survey models.Survey) (*models.Survey, error) {
	return service.repos.CreateSurvey(userID, survey)
}

func (service *SurveyService) GetSurvey(survey models.Survey) (*models.Survey, error) {
	return service.repos.GetSurvey(survey)
}

func (servive *SurveyService) GetSurveys(userID int) (*[]models.Survey, error) {
	return servive.repos.GetSurveys(userID)
}

func (service *SurveyService) DeleteSurvey(survey models.Survey) (*models.Survey, error) {
	return service.repos.DeleteSurvey(survey)
}
