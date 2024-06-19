package handler

import (
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCTX             = "userID"
)

// func (handler *Handler) userIdentity(context *gin.Context) {
// 	cookie, err := context.Cookie("token")
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": "user is not authenticated"})
// 		return
// 	}

// 	userID, err := handler.service.Authorization.ParseToken(cookie)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	context.Set(userCTX, userID)
// }

// func (handler *Handler) _userIdentity(context *gin.Context) {
// 	type inputToken struct {
// 		token string `json:"token" binding:"required"`
// 	}
// 	var input inputToken

// 	if err := context.BindJSON(&input); err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": "user is not authenticated"})
// 		return
// 	}

// 	userID, err := handler.service.Authorization.ParseToken(input.token)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	context.Set(userCTX, userID)
// }

// func getUserID(context *gin.Context) (int, error) {
// 	id, ok := context.Get(userCTX)
// 	if !ok {
// 		return 0, errors.New("user id is not found")
// 	}

// 	return id.(int), nil
// }

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
