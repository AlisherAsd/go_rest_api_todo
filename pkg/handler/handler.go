package handler

import (
	"github.com/gin-gonic/gin"
)

// Структура для обработки запросов
// Содержит методы для обработки запросов которые определены в auth.go, list.go, item.go
type Handler struct {}

// Инициализация маршрутов
// Параметр h *Handler - указатель на структуру Handler
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// Роуты для аутентификации
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	// Роуты для работы с задачами
	// api/
	api := router.Group("/api")
	{
		// api/lists
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			// api/lists/:id/items
			items := lists.Group("/:id/items")
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