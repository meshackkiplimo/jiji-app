package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hlo, World!")
	app := fiber.New()

	type Todo struct {
		ID        int    `json:"id"`
		Body      string `json:"body"`
		completed bool   `json:"completed"`
	}
	todos := []Todo{}

	//routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	app.Post("/api/create", func(c *fiber.Ctx) error {
		todo := &Todo{}
		if err := c.BodyParser(todo); err != nil {
			return err

		}
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Body is required",
			})
		}
		todo.ID = len(todos) + 1
		todos = append(todos, *todo)
		return c.Status(201).JSON(todo)

	})

	log.Fatal(app.Listen(":3000"))
}
