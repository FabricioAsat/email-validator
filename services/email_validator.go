package services

import (
	"regexp"

	"github.com/email-verificator/model"
	"github.com/gofiber/fiber/v2"
)

func EmailValidator(c *fiber.Ctx) error {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	email := new(model.EmailStruct)

	if err := c.BodyParser(email); err != nil || email.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Body error"})
	}

	isValid := emailRegex.MatchString(email.Email)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"isValid": isValid})
}
