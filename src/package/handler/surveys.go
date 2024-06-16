package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (Handler *Handler) getSurveys(context *gin.Context) {
}

func (handler *Handler) createSurvey(context *gin.Context) {
	id, _ := context.Get(userCTX)
	context.JSON(http.StatusOK, gin.H{"id": id})
}

func (handler *Handler) getSurvey(context *gin.Context) {

}
