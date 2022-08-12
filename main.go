package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"romijulianto.io/todoapprj/database"
	"romijulianto.io/todoapprj/models"
)

// Initialize the database connection
func initDatabase() {
	var err error
	dsn := "host=127.0.0.1 user=postgres password=romijulianto dbname=todoapprj port=5432 sslmode=disable"
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connected to database")
	database.DBConn.AutoMigrate(&models.Todo{})
	fmt.Println("Migrated Database")
}

// Routes todo routes to the todo controller
func setupRoutes(app *fiber.App) {
	app.Get("/todos", models.GetTodos)
	app.Get("/todos/:id", models.GetTodoById)
	app.Post("/todos", models.CreateTodo)
	app.Put("/todos/:id", models.UpdateTodo)
	app.Delete("/todos/:id", models.DeleteTodo)
}

// Initialize the server
func main() {
	app := fiber.New()
	initDatabase()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
	setupRoutes(app)
	app.Listen(":3000")
}
