package services

import (
	"github.com/kwhitaker/go-htmx-templ/models"
	"gorm.io/gorm"
)

type TodoService struct {
	db *gorm.DB
}

type TodoServiceInterface interface {
	All() ([]models.Todo, error)
	Get(id uint) (models.Todo, error)
	Create(todo models.Todo) (models.Todo, error)
	Update(todo models.Todo) (models.Todo, error)
	Delete(id uint) error
}

func NewTodoService(db *gorm.DB) *TodoService {
	return &TodoService{db}
}

func (s *TodoService) All() ([]models.Todo, error) {
	var todos []models.Todo
	result := s.db.Find(&todos)

	return todos, result.Error
}

func (s *TodoService) Get(id uint) (models.Todo, error) {
	var todo models.Todo
	result := s.db.First(&todo, id)

	return todo, result.Error
}

func (s *TodoService) Create(todo models.Todo) (models.Todo, error) {
	result := s.db.Create(&todo)

	return todo, result.Error
}

func (s *TodoService) Update(todo models.Todo) (models.Todo, error) {
	result := s.db.Save(&todo)

	return todo, result.Error
}

func (s *TodoService) Delete(id uint) error {
	result := s.db.Delete(&models.Todo{}, id)

	return result.Error
}
