package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func HelloWorld() string {
	return "Hello World!"
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		return c.SendString(HelloWorld())
	})
	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Error starting Fiber server:", err)
	}
}
