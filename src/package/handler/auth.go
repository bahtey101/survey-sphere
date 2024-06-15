package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) signUp(context *gin.Context) {
	type CreateUserInput struct {
		Login    string `json:"login" binding:"required"`
		Passeord string `json:"password" binding:"required"`
	}
	var input CreateUserInput

	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (handler *Handler) signIn(context *gin.Context) {

}
