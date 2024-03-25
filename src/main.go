package main

import (
	"fmt"
	"src/database"
	//"src/handlers"
	//"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	user, err := database.GetUserByLogin("user1")
	if err != nil {
		panic("Failed to get user")
	}
	fmt.Println("USER.ID ", user.ID)

	//user1 := database.GetUserByLogin("user1")

	//surveys := database.GetSurveysByCreator(user1)
	//fmt.Println(surveys)
	//database.GetAllAnswers(surveys[0])
	//questions := database.GetQuestionsBySurvey(surveys[0])
	//fmt.Println(questions)
	//fmt.Println(database.GetPassBySU(surveys[1], user2))
	//router := gin.Default()

	//router.POST("/registration", handlers.Registration)

	//router.Run()
}
