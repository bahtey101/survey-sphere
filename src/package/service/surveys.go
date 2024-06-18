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

func (service *SurveyService) CreateSurvey(survey models.Survey) (*models.Survey, error) {
	return service.repos.CreateSurvey(survey)
}

func (service *SurveyService) GetSurvey(survey models.Survey) (*models.Survey, error) {
	return service.repos.GetSurvey(survey)
}

func (servive *SurveyService) GetSurveys(survey models.Survey) (*[]models.Survey, error) {
	return servive.repos.GetSurveys(survey)
}

func (service *SurveyService) DeleteSurvey(survey models.Survey) (*models.Survey, error) {
	return service.repos.DeleteSurvey(survey)
}

func (service *SurveyService) GetSurveyAnswers(survey models.Survey) (*[]models.Answer, error) {
	return service.repos.GetSurveyAnswers(survey)
}

func (service *SurveyService) GetAllSurveys() (*[]models.Survey, error) {
	return service.repos.GetAllSurveys()
}
