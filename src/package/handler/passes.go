package handler

import (
	"net/http"
	"src/models"

	"github.com/gin-gonic/gin"
)

type PassInput struct {
	Token    string `json:"token" binding:"required"`
	SurveyID uint32 `json:"survey_id"`
}

func (handler *Handler) createPass(context *gin.Context) {
	var input PassInput

	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	respondentID, err := handler.service.Authorization.ParseToken(input.Token)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pass, err := handler.service.CreatePass(models.Pass{
		RespondentID: uint32(respondentID),
		SurveyID:     input.SurveyID,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"pass": pass})
}

func (handler *Handler) getPasses(context *gin.Context) {
	var input PassInput
	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := handler.service.Authorization.ParseToken(input.Token)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passes, err := handler.service.GetPasses(models.Pass{SurveyID: input.SurveyID})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"passes": passes})
}
