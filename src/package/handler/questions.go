package handler

import (
	"net/http"
	"src/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QuestionInput struct {
	Type         models.QuestionType `json:"type" binding:"required"`
	QuestionText string              `json:"question_text" binding:"required"`
}

func (handler *Handler) createQuestion(context *gin.Context) {
	var input QuestionInput
	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	surveyID, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid survey id param"})
		return
	}

	question, err := handler.service.CreateQuestion(models.Question{
		SurveyID:     uint32(surveyID),
		Type:         input.Type,
		QuestionText: input.QuestionText,
	})
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"question": question})
}

func (handler *Handler) getQuestion(context *gin.Context) {
	surveyID, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid survey id param"})
		return
	}

	number, err := strconv.ParseUint(context.Param("question_id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid question id param"})
		return
	}

	question, err := handler.service.Questions.GetQuestion(models.Question{SurveyID: uint32(surveyID), Number: int32(number)})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"question": question})
}

func (handler *Handler) getQuestions(context *gin.Context) {
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

	context.JSON(http.StatusOK, gin.H{"question": question})
}

func (handler *Handler) deleteQuestion(context *gin.Context) {
	surveyID, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid survey id param"})
		return
	}

	number, err := strconv.ParseUint(context.Param("question_id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid question id param"})
		return
	}

	question, err := handler.service.Questions.DeleteQuestion(models.Question{SurveyID: uint32(surveyID), Number: int32(number)})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"question": question})
}
