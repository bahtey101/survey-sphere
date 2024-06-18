package handler

import (
	"net/http"
	"src/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AdminInput struct {
	Token string `json:"token" binding:"required"`
}

func (handler *Handler) getUsers(context *gin.Context) {
	var input AdminInput
	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := handler.service.Authorization.ParseToken(input.Token)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isAdmin := handler.service.Authorization.CheckRole(models.User{ID: uint32(userID)})
	if !isAdmin {
		context.JSON(http.StatusBadRequest, gin.H{"error": "not enough privileges"})
		return
	}

	users, err := handler.service.GetUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"users": users})
}

func (handler *Handler) getAllSurveys(context *gin.Context) {
	type result struct {
		ID           uint32    `json:"survey_id"`
		Topic        string    `json:"topic"`
		Login        string    `json:"login"`
		CreationTime time.Time `json:"creation_time"`
	}

	var input AdminInput
	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := handler.service.Authorization.ParseToken(input.Token)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isAdmin := handler.service.Authorization.CheckRole(models.User{ID: uint32(userID)})
	if !isAdmin {
		context.JSON(http.StatusBadRequest, gin.H{"error": "not enough privileges"})
		return
	}

	surveys, err := handler.service.GetAllSurveys()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var results []result
	for _, survey := range *surveys {
		user, err := handler.service.GetUser(models.User{ID: survey.CreatorID})
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "get user error"})
			return
		}
		results = append(results, result{
			ID:           survey.ID,
			Topic:        survey.Topic,
			Login:        user.Login,
			CreationTime: survey.CreationTime,
		})
	}

	context.JSON(http.StatusOK, gin.H{"surveys": results})
}

func (handler *Handler) deleteUser(context *gin.Context) {
	var input AdminInput
	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := handler.service.Authorization.ParseToken(input.Token)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isAdmin := handler.service.Authorization.CheckRole(models.User{ID: uint32(userID)})
	if !isAdmin {
		context.JSON(http.StatusBadRequest, gin.H{"error": "not enough privileges"})
		return
	}

	id, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := handler.service.DeleteUser(models.User{ID: uint32(id)})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted_user": user})
}
