package controllers

import (
	"go-personal-page/internals/config"
	"go-personal-page/internals/models"
	"go-personal-page/views"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetTodos(c echo.Context) error {
	var todos []models.Todo
	result := config.DB.Find(&todos)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Failed to fetch todos")
	}

	viewTodos := make([]*views.Todo, len(todos))
	for i, todo := range todos {
		viewTodos[i] = &views.Todo{
			Id:          strconv.FormatUint(uint64(todo.ID), 10),
			Description: todo.Description,
		}
	}

	return views.Index(viewTodos).Render(c.Request().Context(), c.Response().Writer)
}

func CreateTodo(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		return c.String(http.StatusBadRequest, "Failed to parse form")
	}

	description := c.Request().Form.Get("description")
	if strings.TrimSpace(description) == "" {
		return c.String(http.StatusBadRequest, "Description cannot be empty")
	}

	todo := models.Todo{
		Description: description,
	}

	result := config.DB.Create(&todo)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Failed to create todo")
	}

	return c.NoContent(http.StatusCreated)
}

func DeleteTodo(c echo.Context) error {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid todo ID")
	}

	result := config.DB.Delete(&models.Todo{}, id)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete todo")
	}

	if result.RowsAffected == 0 {
		return c.String(http.StatusNotFound, "Todo not found")
	}

	return c.NoContent(http.StatusOK)
}
