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
		{
			auth.POST("/sign-up", handler.signUp)
			auth.POST("/sign-in", handler.signIn)
		}
	}

	api := router.Group("/api")
	{
		surveys := api.Group("/surveys")
		{
			surveys.POST("/", handler.getSurveys)
			surveys.POST("/new", handler.createSurvey)
			surveys.POST("/get", handler.getSurveyPasses)

			questions := surveys.Group("/:id")
			{
				questions.POST("/get", handler.getSurveyPasses)
				questions.POST("/questions", handler.getQuestions)
				questions.POST("/new", handler.createQuestion)
			}
		}

		passes := router.Group("/passes")
		{
			passes.POST("/", handler.getPasses)
			passes.POST("/new", handler.createPass)

			answers := passes.Group("/:id")
			{
				answers.POST("/", handler.getAnswers)
				answers.POST("/new", handler.createAnswer)
			}
		}

		admin := router.Group("/admin")
		{
			admin.POST("/users", handler.getUsers)
			admin.POST("/surveys", handler.getAllSurveys)
		}
	}

	return router
}
