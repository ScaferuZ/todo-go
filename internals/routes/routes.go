package routes

import (
	"go-personal-page/internals/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoute(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", controllers.GetTodos)
	e.POST("/", controllers.CreateTodo)
	e.DELETE("/:id", controllers.DeleteTodo)
}
