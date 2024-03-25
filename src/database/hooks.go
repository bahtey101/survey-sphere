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

func (pass *Pass) BeforeDelete(dbase *gorm.DB) error {
	return DeleteAnswersFromPass(pass.ID)
}
