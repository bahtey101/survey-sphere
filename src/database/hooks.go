package database

import (
	"gorm.io/gorm"
)

func (user *User) BeforeDelete(dbase *gorm.DB) (err error) {
	err = DeleteAllUserSurveys(user.ID)
	if err != nil {
		return err
	}

	err = DeleteAllUserPasses(user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (survey *Survey) BeforeDelete(dbase *gorm.DB) (err error) {
	err = DeleteQuestionsFromSurvey(survey.ID)
	if err != nil {
		return err
	}

	err = DeleteAllSurveyPasses(survey.ID)
	if err != nil {
		return err
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

func (pass *Pass) BeforeDelete(dbase *gorm.DB) error {
	return DeleteAnswersFromPass(pass.ID)
}
