package handler

import (
	"github.com/gin-gonic/gin"
	"test/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp) // регистрация
		auth.POST("sign-in", h.signIn) // вход
	}

	api := router.Group("/api", h.userIdentity)
	{
		users := api.Group("/users")
		{
			users.GET("/:id/status", h.getUserInfo)     // вся инфо по юзеру
			users.GET("/leaderboard", h.getLeaderBoard) // топ юзеров с самым большим балансом

			users.POST("/:id/task/complete", h.completeTask) // выполнение задания
			users.POST("/:id/referrer", h.insertReferrer)    // ввод рефералки

			users.GET("/:id/task/list", h.getAllTasksUser) // список заданий, выполненных юзером

		}

		tasks := api.Group("/tasks")
		{
			tasks.GET("/list", h.getAllTasks) // список всех заданий
			tasks.POST("/add", h.addTask)     //
		}
	}

	return router
}
