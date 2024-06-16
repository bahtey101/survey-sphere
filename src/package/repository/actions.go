package repository

import (
	"src/models"

	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB(database *gorm.DB) *gorm.DB {
	DB = database
	return DB
}

// CHECK

func IsExistUser(user *models.User) bool {
	var count int64 = 0
	DB.Find(&user, models.User{Login: user.Login}).Count(&count)
	return count > 0
}

// CREATE

func CreateUser(user *models.User) (*models.User, error) {
	var err error = nil
	if !IsExistUser(user) {
		err = DB.Create(&user).Error
	}
	return GetUser(user), err
}

func CreateSurvey(survey *models.Survey) (*models.Survey, error) {
	err := DB.Create(&survey).Error
	if err != nil {
		return nil, err
	}
	return GetSurvey(survey), nil
}

func CreateQuestion(question *models.Question) (*models.Question, error) {
	err := DB.Create(&question).Error
	if err != nil {
		return nil, err
	}
	return GetQuestion(question), nil
}

func CreatePass(pass *models.Pass) (*models.Pass, error) {
	err := DB.Create(&pass).Error
	if err != nil {
		return nil, err
	}
	return GetPass(pass), nil
}

func CreateAnswer(answer *models.Answer) (*models.Answer, error) {
	err := DB.Create(&answer).Error
	if err != nil {
		return nil, err
	}
	return GetAnswer(answer), nil
}

// GET

func GetUser(user *models.User) *models.User {
	err := DB.Where(&user).First(&user).Error
	if err != nil {
		return nil
	}
	return user
}

func GetSurvey(survey *models.Survey) *models.Survey {
	err := DB.Where(&survey).First(&survey).Error
	if err != nil {
		return nil
	}
	return survey
}

func GetQuestion(question *models.Question) *models.Question {
	err := DB.Where(&question).First(&question).Error
	if err != nil {
		return nil
	}
	return question
}

func GetPass(pass *models.Pass) *models.Pass {
	err := DB.Where(&pass).First(&pass).Error
	if err != nil {
		return nil
	}
	return pass
}

func GetAnswer(answer *models.Answer) *models.Answer {
	err := DB.Where(&answer).First(&answer).Error
	if err != nil {
		return nil
	}
	return answer
}

func GetUsers() *[]models.User {
	var users = []models.User{}
	if err := DB.Find(&users).Error; err != nil {
		return nil
	}
	return &users
}

// DELETE

func DeleteUser(user *models.User) error {
	return DB.Where("login = ?", user.Login).Delete(&models.User{}).Error
}

func DeleteSurvey(survey *models.Survey) error {
	return DB.Where("id = ?", survey.ID).Delete(&models.Survey{}).Error
}
