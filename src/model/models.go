package model

import "time"

type UserRole uint8

const (
	Respondent UserRole = iota
	Admin
)

type QuestionType uint8

const (
	WithText QuestionType = iota
	WithOption
)

type User struct {
	ID       uint16
	Login    string
	Password string
	Role     UserRole
}

type Survey struct {
	ID        uint32
	CreatorID uint16
	Topic     string
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
	Timestamp    time.Time
}

type Answer struct {
	PassID         uint32
	QuestionNumber uint16
	Text           string
}
