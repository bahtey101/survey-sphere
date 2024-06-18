package handler

import (
	"net/http"
	"src/models"
	"strconv"

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

	context.JSON(http.StatusOK, gin.H{"surveys": surveys})
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
