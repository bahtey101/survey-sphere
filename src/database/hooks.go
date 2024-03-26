package database

import (
	// "fmt"

	"gorm.io/gorm"
)

func (question *Question) BeforeCreate(dbase *gorm.DB) (err error) {
	var result int32
	row := GetDataBase().Table(QuestionTable).Where("survey_id = ?", question.SurveyID).Select("MAX(number)").Row()
	err = row.Scan(&result)
	question.Number = result + 1
	return nil
}
