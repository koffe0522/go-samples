package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func helloworld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!!!!")
}

func main() {
	app := fiber.New()

	app.Get("/", helloworld)

	log.Fatal(app.Listen(":3000"))
}
