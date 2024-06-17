package handler

import "github.com/gin-gonic/gin"

type AnswerInput struct {
	QuestionNumber int32  `json:"question_number" binding:"required"`
	AnswerText     string `json:"answer" binding:"required"`
}

func (handler *Handler) createAnswer(context *gin.Context) {

}

func (handler *Handler) getAnswers(context *gin.Context) {

}
