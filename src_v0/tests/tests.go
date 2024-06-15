package tests

import (
	"fmt"
	"src/database"
)

func TEST_1() {
	// Create users
	user1, err := database.CreateUser(&database.User{Login: "user1", Password: "123"})
	if err != nil {
		fmt.Println("UE1")
	}
	fmt.Println("USER1: ", user1)

	user2, err := database.CreateUser(&database.User{Login: "user2", Password: "123"})
	if err != nil {
		fmt.Println("UE2")
	}
	fmt.Println("USER2: ", user2)

	// Create Surveys
	survey1, err := database.CreateSurvey(&database.Survey{CreatorID: user1.ID, Topic: "Nikita_OS"})
	if err != nil {
		fmt.Println("SE1")
	}
	fmt.Println("SURVEY1: ", survey1)

	survey2, err := database.CreateSurvey(&database.Survey{CreatorID: user1.ID, Topic: "Pon_OS"})
	if err != nil {
		fmt.Println("SE2")
	}
	fmt.Println("SURVEY2: ", survey2)

	// Create Questions
	question1, err := database.CreateQuestion(&database.Question{SurveyID: survey1.ID, Type: database.WithText, QuestionText: "ALO?"})
	if err != nil {
		fmt.Println("QE1")
	}
	fmt.Println("QUESTION1: ", question1)

	question2, err := database.CreateQuestion(&database.Question{SurveyID: survey1.ID, Type: database.WithText, QuestionText: "Is it nikitos?"})
	if err != nil {
		fmt.Println("QE2")
	}
	fmt.Println("QUESTION2: ", question2)

	question3, err := database.CreateQuestion(&database.Question{SurveyID: survey2.ID, Type: database.WithText, QuestionText: "HUYEM PO LBU NE DALO?"})
	if err != nil {
		fmt.Println("QE3")
	}
	fmt.Println("QUESTION3: ", question3)

	// Create Passes

	pass1, err := database.CreatePass(&database.Pass{SurveyID: survey1.ID, RespondentID: user1.ID})
	if err != nil {
		fmt.Println("PE1")
	}
	fmt.Println("PASS1: ", pass1)

	pass2, err := database.CreatePass(&database.Pass{SurveyID: survey1.ID, RespondentID: user2.ID})
	if err != nil {
		fmt.Println("PE2")
	}
	fmt.Println("PASS2: ", pass2)

	pass3, err := database.CreatePass(&database.Pass{SurveyID: survey2.ID, RespondentID: user2.ID})
	if err != nil {
		fmt.Println("PE3")
	}
	fmt.Println("PASS3: ", pass3)

	// Create Answers

	answer1, err := database.CreateAnswer(&database.Answer{PassID: pass1.ID, QuestionNumber: 1, AnswerText: "DA, AL0"})
	if err != nil {
		fmt.Println("AE1")
	}
	fmt.Println("ANSWER1: ", answer1)

	answer2, err := database.CreateAnswer(&database.Answer{PassID: pass1.ID, QuestionNumber: 2, AnswerText: "NET BLYAT!"})
	if err != nil {
		fmt.Println("AE2")
	}
	fmt.Println("ANSWER2: ", answer2)

	answer3, err := database.CreateAnswer(&database.Answer{PassID: pass2.ID, QuestionNumber: 1, AnswerText: "SLUSHAYU?"})
	if err != nil {
		fmt.Println("AE3")
	}
	fmt.Println("ANSWER1: ", answer3)

	answer4, err := database.CreateAnswer(&database.Answer{PassID: pass2.ID, QuestionNumber: 2, AnswerText: "POSHEL NAHUY, pidor."})
	if err != nil {
		fmt.Println("AE4")
	}
	fmt.Println("ANSWER1: ", answer4)

	answer5, err := database.CreateAnswer(&database.Answer{PassID: pass3.ID, QuestionNumber: 1, AnswerText: "DA, DAL0"})
	if err != nil {
		fmt.Println("AE5")
	}
	fmt.Println("ANSWER1: ", answer5)

	database.DeleteUser(user1)
	database.DeleteUser(user2)
}

/*
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
	var user2 database.User
	var first_topic = "Test Survey 1"
	var second_topic = "Test Survey 2"
	database.GetDataBase().Table(database.UserTable).Where("login = ?", "user1").First(&user1)
	database.GetDataBase().Table(database.UserTable).Where("login = ?", "user2").First(&user2)
	if !IsExistSurvey(first_topic, user1.ID) {
		database.CreateSurvey(user1.ID, first_topic)
	}
	if !IsExistSurvey(second_topic, user1.ID) {
		database.CreateSurvey(user1.ID, second_topic)
	}

	// Create quesions for Test Survey 1
	var test_survey_1 database.Survey
	database.GetDataBase().Where("creator_id = ? AND topic = ?", user1.ID, first_topic).First(&test_survey_1)
	if !IsExistQuestion(test_survey_1.ID, 1) {
		database.CreateQuestion(test_survey_1.ID, "What is your name?", "with_text")
	}
	if !IsExistQuestion(test_survey_1.ID, 2) {
		database.CreateQuestion(test_survey_1.ID, "What does the fox say?", "with_text")
	}
	if !IsExistQuestion(test_survey_1.ID, 3) {
		database.CreateQuestion(test_survey_1.ID, "Are you male or female?", "with_text")
	}

	// Create quesions for Test Survey 2
	var test_survey_2 database.Survey
	database.GetDataBase().Where("creator_id = ? AND topic = ?", user1.ID, second_topic).First(&test_survey_2)
	if !IsExistQuestion(test_survey_2.ID, 1) {
		database.CreateQuestion(test_survey_2.ID, "What is your name?", "with_text")
	}
	if !IsExistQuestion(test_survey_2.ID, 2) {
		database.CreateQuestion(test_survey_2.ID, "What does the fox say?", "with_text")
	}
	if !IsExistQuestion(test_survey_2.ID, 3) {
		database.CreateQuestion(test_survey_2.ID, "Are you male or female?", "with_text")
	}

	// Create passes for Test Survey 1
	if !IsExistPass(test_survey_1.ID, user2.ID) {
		database.CreatePass(test_survey_1.ID, user2.ID)
	}
	if !IsExistPass(test_survey_1.ID, user1.ID) {
		database.CreatePass(test_survey_1.ID, user1.ID)
	}

	// Create passes for Test Survey 2
	if !IsExistPass(test_survey_2.ID, user2.ID) {
		database.CreatePass(test_survey_2.ID, user2.ID)
	}
	if !IsExistPass(test_survey_2.ID, user1.ID) {
		database.CreatePass(test_survey_2.ID, user1.ID)
	}

	// Create answers for Test Survey 1
	// First pass
	var passTS1U2 database.Pass
	database.GetDataBase().Where("survey_id = ? AND respondent_id = ?", test_survey_1.ID, user2.ID).First(&passTS1U2)
	if !IsExistAnswer(passTS1U2.ID, passTS1U2.SurveyID, 1) {
		database.CreateAnswer(passTS1U2.ID, passTS1U2.SurveyID, 1, "Dimon")
	}
	if !IsExistAnswer(passTS1U2.ID, passTS1U2.SurveyID, 2) {
		database.CreateAnswer(passTS1U2.ID, passTS1U2.SurveyID, 2, "meow")
	}
	if !IsExistAnswer(passTS1U2.ID, passTS1U2.SurveyID, 3) {
		database.CreateAnswer(passTS1U2.ID, passTS1U2.SurveyID, 3, "male")
	}
	// Second pass
	var passTS1U1 database.Pass
	database.GetDataBase().Where("survey_id = ? AND respondent_id = ?", test_survey_1.ID, user1.ID).First(&passTS1U1)
	if !IsExistAnswer(passTS1U1.ID, passTS1U1.SurveyID, 1) {
		database.CreateAnswer(passTS1U1.ID, passTS1U1.SurveyID, 1, "Nikita")
	}
	if !IsExistAnswer(passTS1U1.ID, passTS1U1.SurveyID, 2) {
		database.CreateAnswer(passTS1U1.ID, passTS1U1.SurveyID, 2, "Who knows?")
	}
	if !IsExistAnswer(passTS1U1.ID, passTS1U1.SurveyID, 3) {
		database.CreateAnswer(passTS1U1.ID, passTS1U1.SurveyID, 3, "male")
	}

	// Create answers for Test Survey 2
	// First pass
	var passTS2U2 database.Pass
	database.GetDataBase().Where("survey_id = ? AND respondent_id = ?", test_survey_2.ID, user2.ID).First(&passTS2U2)
	if !IsExistAnswer(passTS2U2.ID, passTS2U2.SurveyID, 1) {
		database.CreateAnswer(passTS2U2.ID, passTS2U2.SurveyID, 1, "Kate")
	}
	if !IsExistAnswer(passTS2U2.ID, passTS2U2.SurveyID, 2) {
		database.CreateAnswer(passTS2U2.ID, passTS2U2.SurveyID, 2, "gav")
	}
	if !IsExistAnswer(passTS2U2.ID, passTS2U2.SurveyID, 3) {
		database.CreateAnswer(passTS2U2.ID, passTS2U2.SurveyID, 3, "female")
	}
	// Second pass
	var passTS2U1 database.Pass
	database.GetDataBase().Where("survey_id = ? AND respondent_id = ?", test_survey_2.ID, user1.ID).First(&passTS2U1)
	if !IsExistAnswer(passTS2U1.ID, passTS2U1.SurveyID, 1) {
		database.CreateAnswer(passTS2U1.ID, passTS2U1.SurveyID, 1, "Masha")
	}
	if !IsExistAnswer(passTS2U1.ID, passTS2U1.SurveyID, 2) {
		database.CreateAnswer(passTS2U1.ID, passTS2U1.SurveyID, 2, "moo")
	}
	if !IsExistAnswer(passTS2U1.ID, passTS2U1.SurveyID, 3) {
		database.CreateAnswer(passTS2U1.ID, passTS2U1.SurveyID, 3, "female")
	}
}

func IsExistSurvey(topic string, creator_id uint32) bool {
	var count int64 = 0
	var survey database.Survey

	database.GetDataBase().Find(&survey, database.Survey{Topic: topic, CreatorID: creator_id}).Count(&count)

	return count > 0
}

func IsExistQuestion(survey_id uint32, number int32) bool {
	var count int64 = 0
	var question database.Question

	database.GetDataBase().Find(&question, database.Question{SurveyID: survey_id, Number: number}).Count(&count)

	return count > 0
}

func IsExistPass(survey_id uint32, respondent_id uint32) bool {
	var count int64 = 0
	var pass database.Pass

	database.GetDataBase().Find(&pass, database.Pass{SurveyID: survey_id, RespondentID: respondent_id}).Count(&count)

	return count > 0
}

func IsExistAnswer(pass_id uint32, survey_id uint32, question_number int32) bool {
	var count int64 = 0
	var answer database.Answer

	database.GetDataBase().Find(&answer, database.Answer{PassID: pass_id, SurveyID: survey_id, QuestionNumber: question_number}).Count(&count)

	return count > 0
}*/
