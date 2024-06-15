package controllers

import (
	"net/http"
	"src/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {
	type CreateUserInput struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	var input CreateUserInput
	var err = context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := database.User{Login: input.Login, Password: input.Password}
	database.CreateUser(&user)

	context.JSON(http.StatusOK, gin.H{"data": user})
}

func FindUsers(context *gin.Context) {
	var users []database.User
	database.GetDataBase().Find(&users)
	context.JSON(http.StatusOK, users)
}

func FindUser(context *gin.Context) {
	id, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	var request = database.User{ID: uint32(id)}
	var user = database.GetUser(&request)
	if user == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(context *gin.Context) {
	id, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	var request = database.User{ID: uint32(id)}
	var user = database.GetUser(&request)
	if user == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	database.GetDataBase().Delete(&user)

	context.JSON(http.StatusOK, gin.H{"data": true})
}
