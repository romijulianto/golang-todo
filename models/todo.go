package models

import (
	"github.com/gofiber/fiber/v2"
	"romijulianto.io/todoapprj/database"
)

type Todo struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Task        string `json:"title"`
	Employee    string `json:"employee"`
	Deadline    string `json:"deadline"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// GetTodos returns all todos
func GetTodos(c *fiber.Ctx) error {
	db := database.DBConn
	var todos []Todo
	db.Find(&todos)
	return c.JSON(todos)
}

// CreateTodo creates a new todo
func CreateTodo(c *fiber.Ctx) error {
	db := database.DBConn
	todo := new(Todo)
	err := c.BodyParser(todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Check your input"})
	}
	err = db.Create(&todo).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create todo"})
	}
	return c.JSON(&todo)
}
