package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCTX             = "userID"
)

func (handler *Handler) userIdentity(context *gin.Context) {
	cookie, err := context.Cookie("token")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "user is not authenticated"})
		return
	}

	userID, err := handler.service.Authorization.ParseToken(cookie)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.Set(userCTX, userID)
}

func getUserID(context *gin.Context) (int, error) {
	id, ok := context.Get(userCTX)
	if !ok {
		return 0, errors.New("user id is not found")
	}

	return id.(int), nil
}
