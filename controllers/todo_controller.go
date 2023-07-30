package controllers

import (
	"net/http"
	"todo/models"
	"todo/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate *validator.Validate = validator.New()

func ListAll(context echo.Context) error {

	todos, err := services.ListAll()

	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})

	}

	return context.JSON(http.StatusOK, models.Response{
		Status:  http.StatusOK,
		Message: "Todos has been listed successfully.",
		Data:    todos,
	})
}

func Retrieve(context echo.Context) error {

	var id string = context.Param("id")

	todo, err := services.FindById(id)

	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return context.JSON(http.StatusOK, models.Response{
		Status:  http.StatusOK,
		Message: "Todo has been retrieved successfully.",
		Data:    todo,
	})
}

func Create(context echo.Context) error {

	var todo models.Todo = models.Todo{}

	if err := context.Bind(&todo); err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if err := validate.Struct(todo); err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err := services.Create(&todo)

	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return context.JSON(http.StatusCreated, models.Response{
		Status:  http.StatusCreated,
		Message: "Todo has been created successfully.",
		Data:    todo,
	})

}

func Update(context echo.Context) error {

	var id string = context.Param("id")
	var todo models.Todo

	if err := context.Bind(&todo); err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if err := validate.Struct(todo); err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err := services.Update(id, &todo)

	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return context.JSON(http.StatusOK, models.Response{
		Status:  http.StatusOK,
		Message: "Todo has been updated successfully.",
		Data:    todo,
	})
}

func Delete(context echo.Context) error {
	var id string = context.Param("id")

	todo, err := services.Delete(id)

	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return context.JSON(http.StatusOK, models.Response{
		Status:  http.StatusOK,
		Message: "Todo has been deleted successfully.",
		Data:    todo,
	})
}
