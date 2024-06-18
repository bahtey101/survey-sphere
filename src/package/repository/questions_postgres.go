package repository

import (
	"src/models"

	"gorm.io/gorm"
)

type QuestionPostgres struct {
	DB *gorm.DB
}

func NewQuestionPostgres(DB *gorm.DB) *QuestionPostgres {
	return &QuestionPostgres{DB: DB}
}

func (postgres *QuestionPostgres) CreateQuestion(question models.Question) (*models.Question, error) {
	var count int64
	postgres.DB.Table("questions").Where(models.Question{SurveyID: question.SurveyID}).Count(&count)
	question.Number = int32(count + 1)

	err := postgres.DB.Create(&question).Error
	if err != nil {
		return nil, err
	}

	return &question, nil
}

func (postgres *QuestionPostgres) GetQuestion(question models.Question) (*models.Question, error) {
	err := postgres.DB.Where(&question).First(&question).Error
	if err != nil {
		return nil, err
	}
	return &question, nil
}

func (postgres *QuestionPostgres) GetQuestions(question models.Question) (*[]models.Question, error) {
	var questions []models.Question
	err := postgres.DB.Where("survey_id = ?", question.SurveyID).Find(&questions).Error
	if err != nil {
		return nil, err
	}
	return &questions, nil
}

func (postgres *QuestionPostgres) DeleteQuestion(question models.Question) (*models.Question, error) {
	err := postgres.DB.Where(&question).First(&question).Error
	if err != nil {
		return nil, err
	}

	err = postgres.DB.Where("survey_id = ? AND number = ?", question.SurveyID, question.Number).Delete(&models.Question{}).Error
	if err != nil {
		return nil, err
	}

	return &question, nil
}

func (postgres *QuestionPostgres) CreateQuestions(questions []models.Question) (*[]models.Question, error) {
	err := postgres.DB.Create(&questions).Error
	if err != nil {
		return nil, err
	}
	return &questions, nil
}
