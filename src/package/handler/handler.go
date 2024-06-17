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

	// auth := router.Group("/auth")
	// {
	// 	auth.POST("/sign-up", handler.signUp)
	// 	auth.GET("/sign-up", func(ctx *gin.Context) {})
	// 	auth.POST("/sign-in", handler.signIn)
	// 	auth.GET("/sign-in", func(ctx *gin.Context) {})
	// }

	// //api := router.Group("/api", handler.userIdentity)
	// api := router.Group("/api")
	// {

	// 	surveys := api.Group("/surveys")
	// 	{
	// 		surveys.POST("/", handler.createSurvey)
	// 		surveys.GET("/", handler.getSurveys)
	// 		surveys.GET("/:id", handler.getSurvey)
	// 		surveys.DELETE("/:id", handler.deleteSurvey)

	// 		questions := surveys.Group(":id/questions")
	// 		{
	// 			questions.POST("/", handler.createQuestion)
	// 			questions.GET("/", handler.getQuestions)
	// 			questions.GET("/:question_id", handler.getQuestion)
	// 			questions.DELETE("/:question_id", handler.deleteQuestion)
	// 		}
	// 	}

	// 	passes := api.Group("/passes")
	// 	{
	// 		passes.POST("/", handler.createPass)
	// 		passes.GET("/", handler.getPasses)
	// 		passes.GET("/:id", handler.getPass)

	// 		answers := passes.Group(":id/answers")
	// 		{
	// 			answers.POST("/", handler.createAnswer)
	// 			answers.GET("/", handler.getAnswers)
	// 		}
	// 	}

	// admin := api.Group("/admin")
	// {
	// 	admin.GET("/", handler.)
	// }

	auth := router.Group("/auth")
	{
		{
			auth.POST("/sign-up", handler.signUp)
			auth.POST("/sign-in", handler.signIn)
		}

		api := router.Group("/api")
		{
			surveys := api.Group("/surveys")
			{
				surveys.POST("/", handler.getSurveys)
				surveys.POST("/new", handler.createSurvey)

				questions := surveys.Group("/:id")
				{
					questions.POST("/")
					questions.POST("/new")
				}
			}

			passes := router.Group("/passes")
			{
				passes.POST("/new")

				answers := passes.Group("/:id")
				{
					answers.POST("/")
					answers.POST("/new")
				}
			}

			// admin := router.Group("/admin")
			// {
			// 	admin.POST("/users", handler.getUsers)
			// 	admin.POST("/surveys", handler.getAllSurveys)
			// }
		}
	}

	return router
}
