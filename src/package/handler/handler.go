package handler

import (
	"src/package/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (handler Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handler.signUp)
		auth.GET("/sign-up", func(ctx *gin.Context) {})
		auth.POST("/sign-in", handler.signIn)
	}

	api := router.Group("/api")
	{
		surveys := api.Group("/surveys")
		{
			surveys.POST("/", handler.createSurvey)
			surveys.GET("/", handler.getSurvey)
		}

	}

	return router
}
