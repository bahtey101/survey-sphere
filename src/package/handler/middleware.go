package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCTX             = "userID"
)

func (handler *Handler) userIdentity(context *gin.Context) {
	header := context.GetHeader(authorizationHeader)
	if header == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "empty auth header"})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid auth header"})
		return
	}

	userID, err := handler.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.Set(userCTX, userID)
}
