package database

import (
	"gorm.io/gorm"
)

func (question *Question) BeforeCreate(dbase *gorm.DB) (err error) {
	var count int64
	GetDataBase().Table("questions").Where(Question{SurveyID: question.SurveyID}).Count(&count)
	question.Number = int32(count + 1)
	return err
}

func (answer *Answer) BeforeCreate(dbase *gorm.DB) (err error) {
	answer.SurveyID = GetPass(&Pass{ID: answer.PassID}).SurveyID
	return err
}
