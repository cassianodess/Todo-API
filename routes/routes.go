package routes

import (
	"todo/controllers"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {

	var todoGroup *echo.Group = e.Group("/todos")
	todoGroup.GET("", controllers.ListAll)
	todoGroup.GET("/:id", controllers.Retrieve)
	todoGroup.POST("", controllers.Create)
	todoGroup.PUT("/:id", controllers.Update)
	todoGroup.DELETE("/:id", controllers.Delete)
}