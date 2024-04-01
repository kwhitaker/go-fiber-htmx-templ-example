package handlers

import (
	"errors"
	"strconv"

	"github.com/a-h/templ"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/kwhitaker/go-htmx-templ/models"
	"github.com/kwhitaker/go-htmx-templ/services"
	"github.com/kwhitaker/go-htmx-templ/view"
	"github.com/kwhitaker/go-htmx-templ/view/layout"
	"github.com/kwhitaker/go-htmx-templ/view/partial"
	"gorm.io/gorm"
)

var validate *validator.Validate

type TodoHandler struct {
	service services.TodoServiceInterface
}

type TodoHandlerInterface interface {
	Index(*fiber.Ctx) error
	TodosList(*fiber.Ctx) error
	Create(*fiber.Ctx) error
	Update(*fiber.Ctx) error
	Delete(*fiber.Ctx) error
}

func NewTodoHandler(db *gorm.DB) *TodoHandler {
	svc := services.NewTodoService(db)
	validate = validator.New(validator.WithRequiredStructEnabled())

	return &TodoHandler{
		service: svc,
	}
}

func (h *TodoHandler) Index(c *fiber.Ctx) error {
	todos, err := h.service.All()

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	l := layout.Base(view.Index(todos))

	return adaptor.HTTPHandler(templ.Handler(l))(c)
}

func (h *TodoHandler) TodosList(c *fiber.Ctx) error {
	todos, err := h.service.All()

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return adaptor.HTTPHandler(templ.Handler(partial.Todos(todos)))(c)
}

func (h *TodoHandler) Create(c *fiber.Ctx) error {
	todo := models.Todo{
		Title:     c.FormValue("title"),
		Completed: false,
	}

	err := validate.Struct(&todo)

	if err != nil {
		err = errors.New("Invalid title")
		return adaptor.HTTPHandler(templ.Handler(partial.TodoForm(err)))(c)
	}

	_, err = h.service.Create(todo)

	if err != nil {
		err = errors.New("Failed to create todo")
		adaptor.HTTPHandler(templ.Handler(partial.TodoForm(err)))(c)
		return c.SendStatus(500)
	}

	c.Append("HX-Trigger", "refresh-todos")
	return adaptor.HTTPHandler(templ.Handler(partial.TodoForm(err)))(c)
}

func (h *TodoHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	numId, err := strconv.Atoi(id)

	if err != nil {
		err = errors.New("Invalid ID")
		return adaptor.HTTPHandler(templ.Handler(partial.TodoItem(nil, false, err)))(c)
	}

	existing, err := h.service.Get(uint(numId))

	if err != nil {
		err = errors.New("Todo not found")
		return adaptor.HTTPHandler(templ.Handler(partial.TodoItem(nil, false, err)))(c)
	}

	updated := existing
	updated.Completed = c.FormValue("completed") == "on"
	updated.Title = c.FormValue("title")

	err = validate.Struct(&updated)

	if err != nil {
		err = errors.New("Invalid title")
		return adaptor.HTTPHandler(templ.Handler(partial.TodoItem(&existing, false, err)))(c)
	}

	updated, err = h.service.Update(updated)

	if err != nil {
		err = errors.New("Failed to update todo")
		return adaptor.HTTPHandler(templ.Handler(partial.TodoItem(&updated, false, err)))(c)
	}

	return adaptor.HTTPHandler(templ.Handler(partial.TodoItem(&updated, true, nil)))(c)
}

func (h *TodoHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	numId, err := strconv.Atoi(id)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	err = h.service.Delete(uint(numId))

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	c.Append("HX-Trigger", "refresh-todos")
	return c.Status(200).SendString("")
}
