package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phum-athiphum/tick-clone-go/config"
)

func main() {
	app := fiber.New()
	config.ConnectDatabase()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Listen(":8080")
}
