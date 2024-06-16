package handler

import (
	"net/http"
	"src/models"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) signUp(context *gin.Context) {
	type CreateUserInput struct {
		Login    string `json:"login" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var input CreateUserInput

	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := handler.service.Authorization.CreateUser(models.User{
		Login:    input.Login,
		Password: input.Password,
	})
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"id": id})
}

func (handler *Handler) signIn(context *gin.Context) {

}
