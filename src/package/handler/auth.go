package handler

import (
	"net/http"
	"src/models"

	"github.com/gin-gonic/gin"
)

type UserInput struct {
	Token    string `json:"Token"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (handler *Handler) signUp(context *gin.Context) {
	var input UserInput

	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Token != "" {
		context.JSON(http.StatusOK, gin.H{"error": "already authz"})
		return
	}

	user, err := handler.service.Authorization.CreateUser(models.User{
		Login:    input.Login,
		Password: input.Password,
	})
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"user": user})
}

func (handler *Handler) signIn(context *gin.Context) {
	var input UserInput

	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Token != "" {
		context.JSON(http.StatusOK, gin.H{"error": "already authz"})
		return
	}

	token, err := handler.service.Authorization.GenerateToken(models.User{
		Login:    input.Login,
		Password: input.Password,
	})
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := handler.service.GetUser(models.User{Login: input.Login})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "user get role"})
		return
	}

	// sending cookies
	context.SetCookie("token", token, 60*60*24, "/", "localhost", false, true)

	context.JSON(http.StatusOK, gin.H{
		"token": token,
		"role":  user.Role,
	})
}
