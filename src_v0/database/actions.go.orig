package database

import (
	//"fmt"
	"time"
)

func IsExistUserByLogin(login string) bool {
	var count int64 = 0
	var user User

<<<<<<< HEAD
	GetDataBase().Table(models.UserTable).Find(&user, models.User{Login: login}).Count(&count)
=======
	GetDataBase().First(&user, User{Login: login}).Count(&count)
>>>>>>> b4c35d06b04524dd6a701c7cea5170df4d84d0bd

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

func CreateAnswer(pass_id uint32, survey_id uint32, qnum uint16, answer string) error {
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

func DeleteAllUserSurveys(creator_id uint32) error {
	return GetDataBase().Delete(&Survey{}, "creator_id = ?", creator_id).Error
}

func DeleteQuestionsFromSurvey(survey_id uint32) error {
	return GetDataBase().Delete(&Question{}, "survey_id = ?", survey_id).Error
}

func DeleteAllSurveyPasses(survey_id uint32) error {
	return GetDataBase().Delete(&Pass{}, "survey_id = ", survey_id).Error
}

func DeleteAllUserPasses(respondent_id uint32) error {
	return GetDataBase().Delete(&Pass{}, "respondent_id = ?", respondent_id).Error
}

func DeleteAnswersFromPass(pass_id uint32) error {
	return GetDataBase().Delete(&Answer{}, "pass_id = ?", pass_id).Error
}

/*func GetAllUserPasses(respondent_id uint32) ([]models.Pass, error) {
	var passes []models.Pass
	err := GetDataBase().Table(models.PassTable).Find(&passes, models.Pass{RespondentID: respondent_id}).Error
	return passes, err
}

func GetAnswersFromPass(pass_id uint32) ([]models.Answer, error) {
	var answers []models.Answer
	err := GetDataBase().Table(models.AnswerTable).Find(&answers, models.Answer{PassID: pass_id}).Error
	return answers, err
}

func GetAllUserSurveys(creator_id uint32) ([]models.Survey, error) {
	var surveys []models.Survey
	err := GetDataBase().Table(models.SurveyTable).Find(&surveys, models.Survey{CreatorID: creator_id}).Error
	return surveys, err
}

func GetQuestionsFromSurvey(survey_id uint32) ([]models.Question, error) {
	var questions []models.Question
	err := GetDataBase().Table(models.QuestionTable).Find(&questions, models.Question{SurveyID: survey_id}).Error
	return questions, err
}

func DeleteSurveyQuestions(survey_id uint32) error {
	surveys := GetAllUserSurveys()
}

func DeleteRespondentAnswers(respondent_id uint32) error {
	respondent_passes, err := GetAllUserPasses(respondent_id)
	if err != nil {
		return err
	}

	for _, pass := range respondent_passes {
		err = DeleteAnswersFromPass(pass.ID)

		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteAnswersFromPass(pass_id uint32) error {
	return GetDataBase().Delete(&models.Answer{}, "pass_id = ?", pass_id).Error
}*/

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
}

func GetAllAnswers(survey Survey) {
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
}*/
