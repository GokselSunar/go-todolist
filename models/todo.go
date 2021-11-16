package models

import (
	"github.com/GokselSunar/go-todolist/database"
	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Title     string `gorm:"size:255;not null" json:"title"`
	Complated bool   `json:"complated"`
}

func GetTodos(c *fiber.Ctx) error {

	db := database.DBConn
	var todos []Todo
	db.Find(&todos)
	return c.JSON(&todos)

}

func CreateToDo(c *fiber.Ctx) error {
	db := database.DBConn
	todo := new(Todo)
	err := c.BodyParser(todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Check Your Input", "data": err})

	}
	err = db.Create(&todo).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could Not Create Todo", "data": err})
	}

	return c.JSON(&todo)

}
