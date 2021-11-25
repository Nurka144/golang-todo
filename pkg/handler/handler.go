package handler

import (
	"github.com/Nurka144/todo-app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signup)
		auth.POST("/sign-in", h.signin)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllList)
			lists.GET("/:id", h.getList)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
		}
	}
	return router
}
