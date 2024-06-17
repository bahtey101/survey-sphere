package handler

import (
	"net/http"
	"src/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AnswerInput struct {
	QuestionNumber int32  `json:"question_number" binding:"required"`
	AnswerText     string `json:"answer" binding:"required"`
}

func (handler *Handler) createAnswer(context *gin.Context) {
	passID, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid pass id param"})
		return
	}

	var input AnswerInput
	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	answer, err := handler.service.CreateAnswer(models.Answer{
		PassID:         uint32(passID),
		QuestionNumber: input.QuestionNumber,
		AnswerText:     input.AnswerText,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"answer": answer})
}

func (handler *Handler) getAnswers(context *gin.Context) {
	passID, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid pass id param"})
		return
	}

	answers, err := handler.service.GetAnswers(models.Answer{PassID: uint32(passID)})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"answer": answers})
}
