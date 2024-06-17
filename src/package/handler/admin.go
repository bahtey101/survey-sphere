package handler

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type AdminInput struct {
// 	Token string `json:"token" binding:"required"`
// }

// func (handler *Handler) getUsers(context *gin.Context) {
// 	var input AdminInput
// 	if err := context.BindJSON(&input); err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	userID, err := handler.service.Authorization.ParseToken(input.Token)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	users, err := handler.service.GetUsers()
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	context.JSON(http.StatusOK, gin.H{"users": users})
// }

// func (handler *Handler) getAllSurveys(context *gin.Context) {
// 	var input AdminInput
// 	if err := context.BindJSON(&input); err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	userID, err := handler.service.Authorization.ParseToken(input.Token)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	surveys, err := handler.service.GetAllSurveys()
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	context.JSON(http.StatusOK, gin.H{"surveys": surveys})
// }
