package handler

import (
	"net/http"
	"src/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SurveyInput struct {
	ID    uint32 `json:"id"`
	Topic string `json:"topic" binding:"required"`
}

func (handler *Handler) createSurvey(context *gin.Context) {
	userID, err := getUserID(context)
	if err != nil {
		return
	}
	input := models.Survey{CreatorID: userID}

	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	survey, err := handler.service.Surveys.CreateSurvey(input)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"survey": survey})
}

func (handler *Handler) getSurveys(context *gin.Context) {
	id, err := getUserID(context)
	if err != nil {
		return
	}

	surveys, err := handler.service.Surveys.GetSurveys(models.Survey{CreatorID: id})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"surveys": surveys})
}

func (handler *Handler) getSurvey(context *gin.Context) {
	creatorID, err := getUserID(context)
	if err != nil {
		return
	}

	surveyID, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id param"})
		return
	}

	survey := &models.Survey{ID: uint32(surveyID), CreatorID: uint32(creatorID)}
	survey, err = handler.service.Surveys.GetSurvey(*survey)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"surveys": survey})
}

func (handler *Handler) deleteSurvey(context *gin.Context) {
	creatorID, err := getUserID(context)
	if err != nil {
		return
	}

	surveyID, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id param"})
		return
	}

	survey := &models.Survey{ID: uint32(surveyID), CreatorID: uint32(creatorID)}
	survey, err = handler.service.Surveys.DeleteSurvey(*survey)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"survey": survey})
}
