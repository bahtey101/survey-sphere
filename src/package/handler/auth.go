package handler

import (
	"net/http"
	"src/models"
	"time"

	"github.com/gin-gonic/gin"
)

type UserInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (handler *Handler) signUp(context *gin.Context) {
	var input UserInput

	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	token, err := handler.service.Authorization.GenerateToken(models.User{
		Login:    input.Login,
		Password: input.Password,
	})
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// context.JSON(http.StatusOK, gin.H{"token": token})
	// context.SetCookie("token", token, time.Now().Add(time.Hour), "/", "localhost", false, true)
	// Устанавливаем токен в куки
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(5 * time.Minute),
		HttpOnly: true,
	}
	http.SetCookie(context.Writer, cookie)
	huy, _ := context.Cookie("jwt")
	context.JSON(http.StatusOK, huy)

}
