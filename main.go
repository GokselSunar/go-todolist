package main

import (
	"fmt"

	"github.com/GokselSunar/go-todolist/database"
	"github.com/GokselSunar/go-todolist/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func helloworld(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func initDatabase() {

	var err error
	dsn := "host=127.0.0.1 user=gorm password=gorm dbname=goTodoList port=5432"
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")

	}
	fmt.Println("Database Conected")
	database.DBConn.AutoMigrate(&models.Todo{})
	fmt.Println("Migrated DB")
}

func setupRoutes(app *fiber.App) {
	app.Get("/todos", models.GetTodos)
	app.Post("/todos", models.CreateToDo)
	app.Get("/todos/:id", models.GetTodoById)
}

func main() {
	app := fiber.New()
	initDatabase()

	app.Get("/", helloworld)
	setupRoutes(app)
	app.Listen(":8081")

}
