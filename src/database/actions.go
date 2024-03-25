package database

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
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
	var user User
	err = GetDataBase().Table(UserTable).Find(&user, User{Login: login}).Error

	if err != nil {
		fmt.Println("GETUSER ERR")
	}

	return GetDataBase().Table(UserTable).Where("id = ?", user.ID).Delete(&user).Error
}

func DeleteUserPasses(user_id uint32) (err error) {
	var passes []Pass
	err = GetDataBase().Table(PassTable).Find(&passes, Pass{RespondentID: user_id}).Error

	if err != nil {
		fmt.Println("GETPASSES ERR")
	}

	for _, pass := range passes {
		err = GetDataBase().Table(PassTable).Where("id = ?", pass.ID).Delete(&pass).Error

		if err != nil {
			fmt.Println("DELPASS ERR")
		}
	}

	return nil
}

func DeleteAnswersByPass(pass_id uint32) (err error) {
	var answers []Answer
	err = GetDataBase().Table(AnswerTable).Find(&answers, Answer{PassID: pass_id}).Error

	if err != nil {
		fmt.Println("GET_ANS_BY_PASS ERR")
	}

	for _, answer := range answers {
		err = GetDataBase().Table(AnswerTable).Where("pass_id = ?", answer.PassID).Delete(&answer).Error

		if err != nil {
			fmt.Println("DEL_ANS_BY_PASS ERR")
		}
	}

	return nil
}

func DeleteUserSurveys(creator_id uint32) (err error) {
	var surveys []Survey
	err = GetDataBase().Table(SurveyTable).Find(&surveys, Survey{CreatorID: creator_id}).Error
	fmt.Println("SURVEYS: PART#1 GET SURVs")
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		fmt.Println("GET_SURs_BY_UID ERR")
	}
	fmt.Println("SURVEYS: PART#2 DEL SURVs")
	for _, survey := range surveys {
		err = GetDataBase().Table(SurveyTable).Where("creator_id = ?", survey.CreatorID).Delete(&survey).Error

		if err != nil {
			fmt.Println("DEL_SUR_BY_UID ERR")
		}
	}

	return nil
}

func DeleteQuestionsBySurvey(survey_id uint32) (err error) {
	var questions []Question
	err = GetDataBase().Table(QuestionTable).Find(&questions, Question{SurveyID: survey_id}).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		fmt.Println("GET_QUEs_BY_SURV ERR")
	}

	for _, question := range questions {
		err = GetDataBase().Table(QuestionTable).Where("survey_id = ?", question.SurveyID).Delete(&question).Error

		if err != nil {
			fmt.Println("DEL_QUE_BY_ANS ERR")
		}
	}

	return nil
}

func DeleteAnswersByQuestion(survey_id uint32, number int32) (err error) {
	var answers []Answer
	err = GetDataBase().Table(AnswerTable).Find(&answers, Answer{SurveyID: survey_id, QuestionNumber: number}).Error

	if err != nil {
		fmt.Println("GET_ANS_BY_QUE ERR")
	}

	for _, answer := range answers {
		err = GetDataBase().Table(AnswerTable).Where("survey_id = ? AND question_number = ?", answer.SurveyID, answer.QuestionNumber).Delete(&answer).Error

		if err != nil {
			fmt.Println("DEL_ANS_BY_QUE ERR")
		}
	}

	return nil
}

func DeletePassesBySurvey(survey_id uint32) (err error) {
	var passes []Pass
	err = GetDataBase().Table(PassTable).Find(&passes, Pass{SurveyID: survey_id}).Error

	if err != nil {
		fmt.Println("GETPASSES_BY_SURV ERR")
	}

	for _, pass := range passes {
		err = GetDataBase().Table(PassTable).Where("survey_id = ?", pass.SurveyID).Delete(&pass).Error

		if err != nil {
			fmt.Println("DELPASS_BY_SURV ERR")
		}
	}

	return nil
}

/*func DeleteUserByLogin(login string) error {
	var user User
	GetDataBase().Table(UserTable).Find(&user, User{Login: login})
	return GetDataBase().Where("1 = 1").Table(UserTable).Delete(&user).Error

	//return GetDataBase().Table(UserTable).Where("login = ?", login).Delete(&User{}).Error
}

func DeleteAllUserSurveys(creator_id uint32) error {
	// var surveys []Survey
	// GetDataBase().Table(SurveyTable).Find(&surveys, Survey{CreatorID: creator_id})
	// return GetDataBase().Where("1 = 1").Delete(&surveys).Error

	return GetDataBase().Table(SurveyTable).Where("creator_id = ?", creator_id).Delete(&Survey{}).Error
}

func DeleteQuestionsFromSurvey(survey_id uint32) error {
	var questions []Question
	GetDataBase().Table(QuestionTable).Find(&questions, Question{SurveyID: survey_id})
	return GetDataBase().Where("1 = 1").Delete(&questions).Error

	//return GetDataBase().Table(QuestionTable).Where("survey_id = ?", survey_id).Delete(&Question{}).Error
}

func DeleteAllSurveyPasses(survey_id uint32) error {
	// var passes []Pass
	// GetDataBase().Table(PassTable).Find(&passes, Pass{SurveyID: survey_id})
	// return GetDataBase().Where("1 = 1").Delete(&passes).Error

	return GetDataBase().Table(PassTable).Where("survey_id = ?", survey_id).Delete(&Pass{}).Error
}

func DeleteAllUserPasses(respondent_id uint32) error {
	// var passes []Pass
	// GetDataBase().Table(PassTable).Find(&passes, Pass{RespondentID: respondent_id})
	// return GetDataBase().Where("1 = 1").Delete(&passes).Error

	return GetDataBase().Table(PassTable).Where("respondent_id = ?", respondent_id).Delete(&Pass{}).Error
}

func DeleteAnswersByQuestions(survey_id uint32, number int32) error {
	var answers []Answer
	GetDataBase().Table(AnswerTable).Find(&answers, Answer{SurveyID: survey_id})
	return GetDataBase().Where("survey_id = ? AND question_number = ?", survey_id, number).Delete(&answers).Error
}

func DeleteAnswersFromPass(pass_id uint32) error {
	// var answers []Answer
	// GetDataBase().Table(AnswerTable).Find(&answers, Answer{PassID: pass_id})
	// return GetDataBase().Where("pass_id = ?", pass_id).Delete(&answers).Error

	return GetDataBase().Table(AnswerTable).Where("pass_id = ?", pass_id).Delete(&Answer{}).Error
}

func GetAllUserPasses(respondent_id uint32) ([]Pass, error) {
	var passes []Pass
	err := GetDataBase().Table(PassTable).Find(&passes, Pass{RespondentID: respondent_id}).Error
	return passes, err
}*/

/*func GetAnswersFromPass(pass_id uint32) ([]models.Answer, error) {
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
