package handler

import (
	"net/http"
	"src/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QuestionInput struct {
	Token        string `json:"token" binding:"required"`
	Number       int32  `json:"number"`
	QuestionText string `json:"question_text"`
}

func (handler *Handler) createQuestion(context *gin.Context) {
	var input QuestionInput
	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := handler.service.Authorization.ParseToken(input.Token)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	surveyID, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid survey id param"})
		return
	}

	question, err := handler.service.Questions.CreateQuestion(models.Question{
		Number:       input.Number,
		QuestionText: input.QuestionText,
		SurveyID:     uint32(surveyID),
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"survey": question})
}

func (handler *Handler) getQuestions(context *gin.Context) {
	var input QuestionInput
	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := handler.service.Authorization.ParseToken(input.Token)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	surveyID, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid survey id param"})
		return
	}

	question, err := handler.service.Questions.GetQuestions(models.Question{SurveyID: uint32(surveyID)})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"survey": question})
}

func (handler *Handler) createQuestions(context *gin.Context) {
	var input struct {
		Token     string `json:"token" binding:"required"`
		Questions []struct {
			Number       int32  `json:"number"`
			QuestionText string `json:"question_text"`
		} `json:"questions" binding:"required"`
	}

	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := handler.service.Authorization.ParseToken(input.Token)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	surveyID, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid survey id param"})
		return
	}

	var questions []models.Question
	for _, question := range input.Questions {
		questions = append(questions, models.Question{
			SurveyID:     uint32(surveyID),
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
