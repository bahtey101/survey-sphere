package repository

import (
	"src/models"

	"gorm.io/gorm"
)

type AnswerPostgres struct {
	DB *gorm.DB
}

func NewAnwserPosgtres(DB *gorm.DB) *AnswerPostgres {
	return &AnswerPostgres{DB: DB}
}

func (postgres *AnswerPostgres) CreateAnswer(answer models.Answer) (*models.Answer, error) {
	err := postgres.DB.Create(&answer).Error
	if err != nil {
		return nil, err
	}
	return postgres.GetAnswer(answer)
}

func (posgres *AnswerPostgres) GetAnswer(answer models.Answer) (*models.Answer, error) {
	err := posgres.DB.Where(&answer).First(&answer).Error
	if err != nil {
		return nil, err
	}
	return &answer, nil
}

func (postgres *AnswerPostgres) GetAnswers(answer models.Answer) (*[]models.Answer, error) {
	var answers *[]models.Answer
	err := postgres.DB.Where("pass_id = ?", answer.PassID).Find(&answers).Error
	if err != nil {
		return nil, err
	}
	return answers, nil
}
