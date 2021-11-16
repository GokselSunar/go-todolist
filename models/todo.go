package models

import (
	"github.com/GokselSunar/go-todolist/database"
	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        uint   `gorm:"primarkey" json:"id"`
	Title     string `json:"title"`
	Complated bool   `json:"complated"`
}

func GetTodos(c *fiber.Ctx) error {

	db := database.DBConn
	var todos []Todo
	db.Find(&todos)
	return c.JSON(&todos)

}
