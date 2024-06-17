package handler

import (
	"net/http"
	"src/models"

	"github.com/gin-gonic/gin"
)

type PassInput struct {
	SurveyID     uint32 `json:"survey_id" binding:"required"`
	ID           uint32 `json:"id"`
	RespondentID uint32 `json:"respondent_id"`
}

func (handler *Handler) createPass(context *gin.Context) {
	userID, err := getUserID(context)
	if err != nil {
		return
	}
	input := models.Pass{RespondentID: userID}

	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pass, err := handler.service.CreatePass(input)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"pass": pass})
}

func (handler *Handler) getPass(context *gin.Context) {
	var input PassInput
	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pass, err := handler.service.GetPass(models.Pass{
		ID:           input.ID,
		RespondentID: input.RespondentID,
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

	passes, err := handler.service.GetPasses(models.Pass{SurveyID: input.SurveyID})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"passes": passes})
}
