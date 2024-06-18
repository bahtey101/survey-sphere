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

func (handler *Handler) createPassWithAnswers(context *gin.Context) {
	var input struct {
		Token    string `json:"token" binding:"required"`
		SurveyID uint32 `json:"survey_id" binding:"required"`
		Answers  []struct {
			QuestionNumber int32  `json:"question_number"`
			AnswerText     string `json:"answer_text"`
		} `json:"answers" binding:"required"`
	}

	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := handler.service.Authorization.ParseToken(input.Token)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pass, err := handler.service.Passes.CreatePass(models.Pass{
		RespondentID: uint32(userID),
		SurveyID:     input.SurveyID,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var answers []models.Answer
	for _, answer := range input.Answers {
		answers = append(answers, models.Answer{
			PassID:         pass.ID,
			SurveyID:       pass.SurveyID,
			QuestionNumber: answer.QuestionNumber,
			AnswerText:     answer.AnswerText,
		})
	}

	_, err = handler.service.CreateAnswers(answers)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "OK"})
}
