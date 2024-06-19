package repository

import (
	"src/models"

	"gorm.io/gorm"
)

type PassPostrgres struct {
	DB *gorm.DB
}

func NewPassPostrgres(DB *gorm.DB) *PassPostrgres {
	return &PassPostrgres{DB: DB}
}

func (postgres *PassPostrgres) CreatePass(pass models.Pass) (*models.Pass, error) {
	err := postgres.DB.Create(&pass).Error
	if err != nil {
		return nil, err
	}
	return postgres.GetPass(pass)
}

func (postgres *PassPostrgres) GetPass(pass models.Pass) (*models.Pass, error) {
	err := postgres.DB.Where(&pass).First(&pass).Error
	if err != nil {
		return nil, err
	}
	return &pass, nil
}

func (postgres *PassPostrgres) GetPasses(pass models.Pass) (*[]models.Pass, error) {
	var passes []models.Pass
	err := postgres.DB.Where("survey_id = ?", pass.SurveyID).Find(&passes).Error
	if err != nil {
		return nil, err
	}
	return &passes, nil
}

func (postgres *PassPostrgres) GetUserPassCount(user models.User) int {
	var count int64
	err := postgres.DB.Where("respondent_id = ?", user.ID).Find(models.Pass{}).Count(&count).Error
	if err != nil {
		return 0
	}
	return int(count)
}
