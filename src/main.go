package main

import (
	"fmt"
	"src/database"
	//"src/handlers"
	//"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	//user1 := database.GetUserByLogin("user1")
	user2 := database.GetUserByLogin("user2")

	//fmt.Println(user2)
	surveys := database.GetSurveysByCreator(user2)
	fmt.Println(surveys)
	//questions := database.GetQuestionsBySurvey(surveys[0])
	//fmt.Println(questions)
	//fmt.Println(database.GetPassBySU(surveys[1], user2))
	//router := gin.Default()

	//router.POST("/registration", handlers.Registration)

	//router.Run()
}
