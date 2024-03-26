package database

import (
	"time"
)

func IsExistUserByLogin(login string) bool {
	var count int64 = 0
	var user User

	GetDataBase().Table(UserTable).Find(&user, User{Login: login}).Count(&count)

	return count > 0
}

// CREATE

func CreateUser(login string, password string, role UserRole) error {
	user := User{Login: login, Password: password, Role: role}
	return GetDataBase().Table(UserTable).Create(&user).Error
}

func CreateSurvey(creator_id uint32, topic string) error {
	survey := Survey{CreatorID: creator_id, Topic: topic, CreationTime: time.Now()}
	return GetDataBase().Table(SurveyTable).Create(&survey).Error
}

func CreateQuestion(survey_id uint32, question_text string, question_type QuestionType) error {
	question := Question{SurveyID: survey_id, QuestionText: question_text, Type: question_type}
	return GetDataBase().Table(QuestionTable).Create(&question).Error
}

func CreatePass(survey_id uint32, respondent_id uint32) error {
	pass := Pass{SurveyID: survey_id, RespondentID: respondent_id, CreationTime: time.Now()}
	return GetDataBase().Table(PassTable).Create(&pass).Error
}

func CreateAnswer(pass_id uint32, survey_id uint32, qnum int32, answer string) error {
	ans := Answer{PassID: pass_id, SurveyID: survey_id, QuestionNumber: qnum, AnswerText: answer}
	return GetDataBase().Table(AnswerTable).Create(&ans).Error
}

// GET

func GetUserByLogin(login string) (User, error) {
	var user User
	err := GetDataBase().Table(UserTable).Where(User{Login: login}).First(&user).Error
	return user, err
}

// DELETE

func DeleteUserByLogin(login string) (err error) {
	return GetDataBase().Table(UserTable).Where("login = ?", login).Delete(&User{}).Error
}

func DeleteSurveyByID(survey_id uint32) (err error) {
	return GetDataBase().Table(SurveyTable).Where("id = ?", survey_id).Delete(&Survey{}).Error
}
