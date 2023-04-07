package main

import (
	"github.com/email-verificator/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/", services.EmailValidator)

	app.Listen(":3000")
}
