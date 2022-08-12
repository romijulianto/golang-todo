package models

import (
	"github.com/gofiber/fiber/v2"
	"romijulianto.io/todoapprj/database"
)

type Todo struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Task        string `json:"task"`
	Employee    string `json:"employee"`
	Deadline    string `json:"deadline"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
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

// GetTodos returns all todos
func GetTodos(c *fiber.Ctx) error {
	db := database.DBConn
	var todos []Todo
	db.Find(&todos)
	return c.JSON(todos)
}

// GetTodo returns a todo by id
func GetTodoById(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var todo Todo
	err := db.Find(&todo, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Could not find todo"})
	}
	return c.JSON(&todo)
}

// UpdateTodo updates a todo
func UpdateTodo(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var todo Todo
	err := db.Find(&todo, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Could not find todo"})
	}
	err = c.BodyParser(&todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Check your input"})
	}
	err = db.Save(&todo).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update todo"})
	}
	return c.JSON(&todo)
}

// DeleteTodo deletes a todo
func DeleteTodo(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var todo Todo
	err := db.Find(&todo, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Could not find todo"})
	}
	err = db.Delete(&todo).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete todo"})
	}
	return c.JSON(&todo)

}
