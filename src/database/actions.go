package database

import (
	"fmt"
	"src/models"
	"time"
)

func IsExistUserByLogin(login string) bool {
	var count int64 = 0
	var user models.User

	GetDataBase().Table(models.UserTable).Find(&user, models.User{Login: login}).Count(&count)

	return count > 0
}

func CreateUser(login string, password string, role models.UserRole) error {
	user := models.User{Login: login, Password: password, Role: role}
	return GetDataBase().Table(models.UserTable).Create(&user).Error
}

func CreateSurvey(creator_id uint32, topic string) error {
	survey := models.Survey{CreatorID: creator_id, Topic: topic, CreationTime: time.Now()}
	return GetDataBase().Table(models.SurveyTable).Create(&survey).Error
}

func CreateQuestion(survey_id uint32, question_text string, question_type models.QuestionType) error {
	question := models.Question{SurveyID: survey_id, QuestionText: question_text, Type: question_type}
	return GetDataBase().Table(models.QuestionTable).Create(&question).Error
}

func CreatePass(survey_id uint32, respondent_id uint32) error {
	pass := models.Pass{SurveyID: survey_id, RespondentID: respondent_id, CreationTime: time.Now()}
	return GetDataBase().Table(models.PassTable).Create(&pass).Error
}

func CreateAnswer(pass_id uint32, survey_id uint32, qnum uint16, answer string) error {
	ans := models.Answer{PassID: pass_id, SurveyID: survey_id, QuestionNumber: qnum, AnswerText: answer}
	return GetDataBase().Table(models.AnswerTable).Create(&ans).Error
}

func GetUserByLogin(login string) models.User {
	var user models.User
	GetDataBase().Table(models.UserTable).Where(models.User{Login: login}).First(&user)
	return user
}

/*func GetUser(user models.User) models.User {
	GetDataBase().Table(models.UserTable).First(&user)
	return user
}

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
}*/

func GetAllAnswers(survey models.Survey) {
	var questions []models.Question
	GetDataBase().Table(models.QuestionTable).Find(&questions, models.Question{SurveyID: survey.ID})
	var pass []models.Pass
	GetDataBase().Table(models.PassTable).Find(&pass, models.Pass{SurveyID: survey.ID})

	fmt.Println(survey.Topic)
	for i := 0; i < len(questions); i += 1 {
		fmt.Println("Q[", i+1, "]: ", questions[i].QuestionText)
		for j := 0; j < len(pass); j += 1 {
			var user models.User
			GetDataBase().Table(models.UserTable).First(&user, pass[j].RespondentID)
			var answer models.Answer
			GetDataBase().Table(models.AnswerTable).Find(&answer, models.Answer{PassID: pass[j].ID, QuestionNumber: uint16(i + 1)})
			fmt.Println("A[", i+1, "][", j+1, "]: Login:", user.Login, " Text:", answer.AnswerText)
		}
	}
}
