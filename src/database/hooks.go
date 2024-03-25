package database

import (
	"fmt"

	"gorm.io/gorm"
)

func (user *User) BeforeDelete(dbase *gorm.DB) (err error) {
	fmt.Println("PART#1 DEL PASSES")
	fmt.Println("MY USER: ", user.ID)
	//passes
	err = DeleteUserPasses(user.ID)

	if err != nil {
		fmt.Println("DEL_U_PASSES ERR")
	}

	fmt.Println("PART#2 DEL SURVs")
	//surveys
	err = DeleteUserSurveys(user.ID)

	if err != nil {
		fmt.Println("DEL_U_SURVES ERR")
	}

	return nil
}

func (pass *Pass) BeforeDelete(dbase *gorm.DB) (err error) {
	fmt.Println("MY PASS: ", pass.ID)
	//answers
	err = DeleteAnswersByPass(pass.ID)

	if err != nil {
		fmt.Println("DEL_ANSSS_BY_PASS ERR")
	}

	return nil
}

func (survey *Survey) BeforeDelete(dbase *gorm.DB) (err error) {
	fmt.Println("MY SURVEY: ", survey.ID)
	//questions
	err = DeleteQuestionsBySurvey(survey.ID)

	if err != nil {
		fmt.Println("DEL_QUE_BY_SURV ERR")
	}
	//passes
	err = DeletePassesBySurvey(survey.ID)

	if err != nil {
		fmt.Println("DEL_PASs_BY_SURV ERR")
	}

	return nil
}

func (question Question) BeforeDelete(dbase *gorm.DB) (err error) {
	//answers
	err = DeleteAnswersByQuestion(question.SurveyID, question.Number)

	if err != nil {
		fmt.Println("DEL_QUE_BY_SURV ERR")
	}

	return nil
}

func (question *Question) BeforeCreate(dbase *gorm.DB) (err error) {
	var result int32
	row := GetDataBase().Table(QuestionTable).Where("survey_id = ?", question.SurveyID).Select("MAX(number)").Row()
	err = row.Scan(&result)
	question.Number = result + 1
	return nil
}
