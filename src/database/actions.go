package database

import "fmt"

// CHECK

func IsExistUser(user *User) bool {
	var count int64 = 0
	GetDataBase().Find(&user, User{Login: user.Login}).Count(&count)
	return count > 0
}

// CREATE TABLES

func CreateTables() {
	if !GetDataBase().Migrator().HasTable(&User{}) {
		GetDataBase().Migrator().CreateTable(&User{})
	}
	if !GetDataBase().Migrator().HasTable(&Survey{}) {
		GetDataBase().Migrator().CreateTable(&Survey{})
	}
	if !GetDataBase().Migrator().HasTable(&Question{}) {
		GetDataBase().Migrator().CreateTable(&Question{})
	}
	if !GetDataBase().Migrator().HasTable(&Pass{}) {
		GetDataBase().Migrator().CreateTable(&Pass{})
	}
	if !GetDataBase().Migrator().HasTable(&Answer{}) {
		GetDataBase().Migrator().CreateTable(&Answer{})
	}
}

// CREATE

func CreateUser(user *User) (*User, error) {
	var err error = nil
	if !IsExistUser(user) {
		err = GetDataBase().Create(&user).Error
	}
	return GetUser(user), err
}

func CreateSurvey(survey *Survey) (*Survey, error) {
	err := GetDataBase().Create(&survey).Error
	if err != nil {
		return nil, err
	}
	return GetSurvey(survey), nil
}

func CreateQuestion(question *Question) (*Question, error) {
	err := GetDataBase().Create(&question).Error
	if err != nil {
		return nil, err
	}
	return GetQuestion(question), nil
}

func CreatePass(pass *Pass) (*Pass, error) {
	err := GetDataBase().Create(&pass).Error
	if err != nil {
		return nil, err
	}
	return GetPass(pass), nil
}

func CreateAnswer(answer *Answer) (*Answer, error) {
	err := GetDataBase().Create(&answer).Error
	if err != nil {
		return nil, err
	}
	return GetAnswer(answer), nil
}

// GET

func GetUser(user *User) *User {
	err := GetDataBase().Where(&user).First(&user).Error
	if err != nil {
		return nil
	}
	return user
}

func GetSurvey(survey *Survey) *Survey {
	err := GetDataBase().Where(&survey).First(&survey).Error
	if err != nil {
		return nil
	}
	return survey
}

func GetQuestion(question *Question) *Question {
	err := GetDataBase().Where(&question).First(&question).Error
	if err != nil {
		return nil
	}
	return question
}

func GetPass(pass *Pass) *Pass {
	err := GetDataBase().Where(&pass).First(&pass).Error
	if err != nil {
		return nil
	}
	return pass
}

func GetAnswer(answer *Answer) *Answer {
	err := GetDataBase().Where(&answer).First(&answer).Error
	if err != nil {
		return nil
	}
	return answer
}

func GetUsers() *[]User {
	var users = []User{}
	err := GetDataBase().Find(users).Error
	if err != nil {
		return nil
	}
	fmt.Println("Юзвери")
	fmt.Println(users)
	return &users
}

// DELETE

func DeleteUser(user *User) error {
	return GetDataBase().Where("login = ?", user.Login).Delete(&User{}).Error
}

func DeleteSurvey(survey *Survey) error {
	return GetDataBase().Where("id = ?", survey.ID).Delete(&Survey{}).Error
}
