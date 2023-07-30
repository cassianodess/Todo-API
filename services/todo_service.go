package services

import (
	"todo/models"
	"todo/repository"

	"gorm.io/gorm/clause"
)

func ListAll() ([]models.Todo, error) {
	var todos []models.Todo
	result := repository.Context.Find(&todos)
	return todos, result.Error
}

func FindById(id string) (models.Todo, error) {
	var todo models.Todo
	result := repository.Context.First(&todo, "id = ?", id)
	return todo, result.Error
}

func Create(todo *models.Todo) error {
	result := repository.Context.Create(&todo)
	return result.Error
	
}

func Update(id string, todo *models.Todo) error {
	var updatedTodo models.Todo
	result := repository.Context.Find(&updatedTodo, "id = ?", id)

	if result.Error != nil {
		return result.Error
	}
	
	updatedTodo.Title = todo.Title
	updatedTodo.Completed = todo.Completed

	updatedResult := repository.Context.Save(&updatedTodo)
	*todo = updatedTodo

	return updatedResult.Error

}

func Delete(id string) (models.Todo, error) {
	var todo models.Todo
	result := repository.Context.Clauses(clause.Returning{}).Delete(&todo, "id = ?", id)
	return todo, result.Error
}