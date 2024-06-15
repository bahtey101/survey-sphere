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
	ID       uint32   `gorm:"uniqueIndex;primaryKey;"`
	Login    string   `gorm:"unique;not null;"`
	Password string   `gorm:"not null;"`
	Role     UserRole `gorm:"not null;default:respondent"`
	Surveys  []Survey `gorm:"foreignKey:CreatorID;references:ID;constraint:OnDelete:CASCADE;"`
	Passes   []Pass   `gorm:"foreignKey:RespondentID;references:ID;constraint:OnDelete:CASCADE;"`
}

func New(login string, password string) *User {
	return &User{Login: login, Password: password, Role: Respondent}
}

type Survey struct {
	ID           uint32     `gorm:"uniqueIndex;primaryKey;"`
	CreatorID    uint32     `gorm:"references:User.ID;index;"`
	Topic        string     `gorm:"index;unique;"`
	CreationTime time.Time  `gorm:"not null;autoCreateTime:milli;"`
	Questions    []Question `gorm:"foreignKey:SurveyID;references:ID;constraint:OnDelete:CASCADE;"`
}

type Question struct {
	SurveyID     uint32 `gorm:"references:Survey.ID;index;primaryKey;"`
	Number       int32  `gorm:"index;primaryKey;default:1;"`
	Type         QuestionType
	QuestionText string
}

type Pass struct {
	ID           uint32    `gorm:"uniqueIndex;primaryKey;"`
	SurveyID     uint32    `gorm:"references:Survey.ID;index;"`
	RespondentID uint32    `gorm:"references:User.ID;index;"`
	CreationTime time.Time `gorm:"not null;autoCreateTime:milli"`
	Answers      []Answer  `gorm:"foreignKey:PassID;references:ID;constraint:OnDelete:CASCADE;"`
}

type Answer struct {
	PassID         uint32 `gorm:"references:Pass.ID;index;primaryKey;"`
	SurveyID       uint32 `gorm:"references:Survey.ID;index;primaryKey;"`
	QuestionNumber int32  `gorm:"index;primaryKey;"`
	AnswerText     string
}
