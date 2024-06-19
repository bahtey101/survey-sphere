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

	all := router.Group("/", CORSMiddleware())
	{
		auth := all.Group("/auth")
		{
			{
				auth.POST("/sign-up", handler.signUp)
				auth.POST("/sign-in", handler.signIn)
			}
		}

		api := all.Group("/api")
		{

			surveys := api.Group("/surveys")
			{
				surveys.POST("/", handler.getSurveys)
				surveys.POST("/new", handler.createSurveyWithQuestions)
				surveys.POST("/survey/:id", handler.getSurveyWithQuestions)

				questions := surveys.Group("/:id")
				{
					questions.POST("/get", handler.getSurveyAnswers)
					questions.POST("/questions", handler.getSurveyWithQuestions)
				}
			}

			passes := api.Group("/passes")
			{
				passes.POST("/", handler.getPasses)
				passes.POST("/new", handler.createPassWithAnswers)

				answers := passes.Group("/:id")
				{
					answers.POST("/", handler.getAnswers)
				}
			}

			admin := api.Group("/admin")
			{
				admin.POST("/users", handler.getUsers)
				admin.POST("/surveys", handler.getAllSurveys)
				admin.POST("/users/:id", handler.deleteUser)
				admin.POST("/surveys/:id", handler.deleteSurvey)
			}
		}
	}

	return router
}
