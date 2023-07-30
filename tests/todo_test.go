package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"todo/controllers"
	"todo/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	repository, err := models.PostgresDataBase{}.Connect()

	if err != nil {
		log.Fatalf("Fail in connect with database: %v", err)
	}

	db = repository

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestListTodosSuccessfully(t *testing.T) {
	t.Run("Should list todos with status 200 OK", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/todos", nil)
		recorder := httptest.NewRecorder()
		e.NewContext(request, recorder)

		if recorder.Code != http.StatusOK {
			t.Errorf("Expected %d, but got %d", http.StatusOK, recorder.Code)
		}

	})
}

func TestCreateTodoSuccessfully(t *testing.T) {
	t.Run("Should create status 201 Created", func(t *testing.T) {
		e := echo.New()

		todo := models.Todo{
			Title: "Test",
		}

		body, err := json.Marshal(todo)

		if err != nil {
			t.Fatalf("Error in serialize json object: %v", err)
		}

		request := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(body))
		request.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()
		context := e.NewContext(request, recorder)

		controllers.Create(context)

		if recorder.Code != http.StatusCreated {
			t.Errorf("Expected %d, but got %d", http.StatusCreated, recorder.Code)
		}

		db.Delete(&todo, "title = ?", todo.Title)
	})

}
func TestCreateTodoWithInvalidBody(t *testing.T) {
	t.Run("Should not create status 400 BAD REQUEST", func(t *testing.T) {
		e := echo.New()

		todo := models.Todo{
			Title: "",
		}

		body, err := json.Marshal(todo)

		if err != nil {
			t.Fatalf("Error in serialize json object: %v", err)
		}

		request := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(body))
		request.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()
		context := e.NewContext(request, recorder)

		controllers.Create(context)

		if recorder.Code != http.StatusBadRequest {
			t.Errorf("Expected %d, but got %d", http.StatusBadRequest, recorder.Code)
		}
	})
}
