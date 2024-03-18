package models

import "time"

const (
	UserTable     string = "site_user"
	SurveyTable          = "survey"
	QuestionTable        = "question"
	PassTable            = "pass"
	AnswerTable          = "answer"
)

type UserRole string

const (
	Respondent UserRole = "respondent"
	Admin               = "admin"
)

type QuestionType string

const (
	WithText   QuestionType = "with_text"
	WithOption              = "with_option"
)

type User struct {
	ID       uint16
	Login    string
	Password string
	Role     UserRole
}

type Survey struct {
	ID           uint32
	CreatorID    uint16
	Topic        string
	CreationTime time.Time
}

type Question struct {
	SurveyID uint32
	Number   uint16
	Type     string
	Text     string
}

type Pass struct {
	ID           uint32
	SurveyID     uint32
	RespondentID uint16
	CreationTime time.Time
}

type Answer struct {
	PassID         uint32
	QuestionNumber uint16
	Text           string
}
