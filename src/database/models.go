package database

import "time"

const (
	UserTable     string = "site_users"
	SurveyTable   string = "surveys"
	QuestionTable string = "questions"
	PassTable     string = "passes"
	AnswerTable   string = "answers"
)

type UserRole string

const (
	Respondent UserRole = "respondent"
	Admin      UserRole = "admin"
)

type QuestionType string

const (
	WithText   QuestionType = "with_text"
	WithOption QuestionType = "with_option"
)

type User struct {
	ID       uint32
	Login    string
	Password string
	Role     UserRole
}

type Survey struct {
	ID           uint32
	CreatorID    uint32
	Topic        string
	CreationTime time.Time
}

type Question struct {
	SurveyID     uint32
	Number       int32
	Type         QuestionType
	QuestionText string
}

type Pass struct {
	ID           uint32
	SurveyID     uint32
	RespondentID uint32
	CreationTime time.Time
}

type Answer struct {
	PassID         uint32
	SurveyID       uint32
	QuestionNumber int32
	AnswerText     string
}
