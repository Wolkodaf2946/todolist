package handler

import (
	"github.com/Wolkodaf2946/todolist/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine { // инициализация эндпоинтов
	router := gin.New() // инициализация роутера

	// объявляем методы, сгруппировав их по маршрутам

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp) // регистрация
		auth.POST("/sign-in", h.signIn) // авторизация
	}

	api := router.Group("/api") // будет использоваться для эндпоинтов работы с списками и их задачами
	{
		lists := api.Group("/lists") // работа со списками
		{
			lists.POST("/", h.createList)      // создание
			lists.GET("/", h.getAllLists)      // получение всех списков
			lists.GET("/:id", h.getListById)   // получение списка по id
			lists.PUT("/:id", h.updateList)    // редактирование списка
			lists.DELETE("/:id", h.deleteList) // удаление

			items := lists.Group(":id/items") // группа для задач списков
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}
	return router
}
