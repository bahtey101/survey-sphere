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
		auth.GET("/sign-in", func(ctx *gin.Context) {})
	}

	api := router.Group("/api")
	{

		surveys := api.Group("/mysurveys", handler.userIdentity)
		{
			surveys.GET("/", handler.getSurveys)
			surveys.GET("/:id", handler.getSurvey)
			surveys.POST("/new", handler.createSurvey)
		}

		// admin := api.Group("/admin")
		// {
		// 	admin.GET("/", handler.)
		// }

	}

	return router
}
