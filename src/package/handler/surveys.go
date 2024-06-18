package handler

import (
	"net/http"
	"src/models"

	"github.com/gin-gonic/gin"
)

type SurveyInput struct {
	Token string `json:"token" binding:"required"`
	Topic string `json:"topic"`
}

func (handler *Handler) createSurvey(context *gin.Context) {
	var input SurveyInput
	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := handler.service.Authorization.ParseToken(input.Token)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	survey, err := handler.service.Surveys.CreateSurvey(models.Survey{
		CreatorID: uint32(userID),
		Topic:     input.Topic,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"survey": survey})
}

func (handler *Handler) getSurveys(context *gin.Context) {
	var input SurveyInput
	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := handler.service.Authorization.ParseToken(input.Token)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	surveys, err := handler.service.Surveys.GetSurveys(models.Survey{CreatorID: uint32(userID)})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"surveys": surveys})
}

func (handler *Handler) getSurveyPasses(context *gin.Context) {
	var input SurveyInput
	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := handler.service.Authorization.ParseToken(input.Token)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passes, err := handler.service.Surveys.GetSurveyPasses(models.Survey{CreatorID: uint32(userID)})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"passes": passes})
}

func (handler *Handler) createSurveyWithQuestions(context *gin.Context) {
	var input struct {
		Token     string `json:"token" binding:"required"`
		Topic     string `json:"topic" binding:"required"`
		Questions []struct {
			Number       int32  `json:"number"`
			QuestionText string `json:"question_text"`
		} `json:"questions" binding:"required"`
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

	survey, err := handler.service.Surveys.CreateSurvey(models.Survey{
		CreatorID: uint32(userID),
		Topic:     input.Topic,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var questions []models.Question
	for _, question := range input.Questions {
		questions = append(questions, models.Question{
			SurveyID:     uint32(survey.ID),
			Number:       question.Number,
			QuestionText: question.QuestionText,
		})
	}

	_, err = handler.service.CreateQuestions(questions)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "OK"})
}
