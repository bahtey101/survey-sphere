package handler

import "github.com/gin-gonic/gin"

type PassInput struct {
	RespondentID uint32 `json:"respondent_id" binding:"required"`
	SurveyID     uint32 `json:"survey_id" binding:"required"`
}

func (handler *Handler) createPass(context *gin.Context) {

}

func (handler *Handler) getPass(context *gin.Context) {

}

func (handler *Handler) getPasses(context *gin.Context) {

}
