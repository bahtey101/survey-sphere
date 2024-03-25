package tests

import (
	"fmt"
	"src/database"
)

func InitTestValues() {
	// Create site_users
	if !database.IsExistUserByLogin("user1") {
		database.CreateUser("user1", "1234", "respondent")
	}
	if !database.IsExistUserByLogin("user2") {
		database.CreateUser("user2", "4321", "respondent")
	}

	// Create user1's surveys
	var user1 database.User
	var first_topic = "Test Survey 1"
	var second_topic = "Test Survey 2"
	database.GetDataBase().Table(database.UserTable).Where("login = ?", "user1").First(&user1)
	if !IsExistSurvey(first_topic, user1.ID) {
		database.CreateSurvey(user1.ID, first_topic)
	}
	if !IsExistSurvey(second_topic, user1.ID) {
		database.CreateSurvey(user1.ID, second_topic)
	}

	// Create quesions for Test Survey 1
	var test_survey_1 database.Survey
	database.GetDataBase().Where("creator_id = ? AND topic = ?", user1.ID, first_topic).First(&test_survey_1)
	fmt.Println(test_survey_1)
	database.CreateQuestion(test_survey_1.ID, "What is your name?", "with_text")
	// if !IsExistQuestion(test_survey_1.ID, 1) {
	// 	database.CreateQuestion(test_survey_1.ID, "What is your name?", "with_text")
	// }
	// if !IsExistQuestion(test_survey_1.ID, 2) {
	// 	database.CreateQuestion(test_survey_1.ID, "What does the fox say?", "with_text")
	// }
	// if !IsExistQuestion(test_survey_1.ID, 3) {
	// 	database.CreateQuestion(test_survey_1.ID, "Are you male or female?", "with_text")
	// }
}

func IsExistSurvey(topic string, creator_id uint32) bool {
	var count int64 = 0
	var survey database.Survey

	database.GetDataBase().Find(&survey, database.Survey{Topic: topic, CreatorID: creator_id}).Count(&count)

	return count > 0
}

func IsExistQuestion(survey_id uint32, number uint16) bool {
	var count int64 = 0
	var question database.Question

	database.GetDataBase().Find(&question, database.Question{SurveyID: survey_id, Number: number}).Count(&count)

	return count > 0
}
