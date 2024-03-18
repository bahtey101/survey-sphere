package database

import "src/models"

func IsExistUserByLogin(login string) bool {
	var count int64 = 0
	var user models.User

	GetDataBase().First(&user, models.User{Login: login}).Count(&count)

	return count > 0
}

func Add(bean interface{}) error {
	return GetDataBase().Create(bean).Error
}

func GetUserByLogin(login string) models.User {
	var user models.User
	GetDataBase().Table(models.UserTable).Where(models.User{Login: login}).First(&user)
	return user
}

/*func GetUser(user models.User) models.User {
	GetDataBase().Table(models.UserTable).First(&user)
	return user
}*/

func GetSurveysByCreator(creator models.User) []models.Survey {
	var surveys []models.Survey
	GetDataBase().Table(models.SurveyTable).Find(&surveys, models.Survey{CreatorID: creator.ID})
	return surveys
}

func GetQuestionsBySurvey(survey models.Survey) []models.Question {
	var questions []models.Question
	GetDataBase().Table(models.QuestionTable).Find(&questions, models.Question{SurveyID: survey.ID})
	return questions
}

func GetPassBySU(survey models.Survey, user models.User) models.Pass {
	var pass models.Pass
	GetDataBase().Table(models.PassTable).Find(&pass, models.Pass{SurveyID: survey.ID, RespondentID: user.ID})
	return pass
}

func GetAnswersByPass(pass models.Pass) []models.Answer {
	var answers []models.Answer
	GetDataBase().Table(models.AnswerTable).Find(&answers, models.Answer{PassID: pass.ID})
	return answers
}
