package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hlo, World!")
	app := fiber.New()
	log.Fatal(app.Listen(":3000"))
}
