package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/kwhitaker/go-htmx-templ/handlers"
	"github.com/kwhitaker/go-htmx-templ/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// obv you would use an environment variable in a real app
	mode := os.Getenv("MODE")

	log.Println("Running in mode:", mode)

	dbName := "todos.db"

	if mode == "test" {
		dbName = "test.db"
	}

	log.Println("Using database:", dbName)

	db, err := gorm.Open(sqlite.Open(dbName), nil)

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Todo{})

	app := fiber.New()
	app.Static("/", "./static")

	todosHandler := handlers.NewTodoHandler(db)

	app.Get("/", todosHandler.Index)
	app.Get("/todos-list", todosHandler.TodosList)
	app.Post("/todos", todosHandler.Create)
	app.Put("/todos/:id", todosHandler.Update)
	app.Delete("/todos/:id", todosHandler.Delete)

	app.Listen(":3000")
}
