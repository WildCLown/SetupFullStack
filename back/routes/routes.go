package routes

import (
	"go-back/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/items", controllers.GetItems)
		api.POST("/items", controllers.CreateItem)
		api.GET("/items/:id", controllers.GetItem)
		api.PUT("/items/:id", controllers.UpdateItem)
		api.DELETE("/items/:id", controllers.DeleteItem)
	}
}
