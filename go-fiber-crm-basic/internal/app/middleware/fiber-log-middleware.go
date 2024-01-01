package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func LoggingMiddleware(c *fiber.Ctx) error {
	log.Println("Request:", c.Method(), c.Path())
	err := c.Next()
	if err != nil {
		return err
	}
	log.Println("Response:", c.Response().StatusCode(), "Body:", string(c.Response().Body()))
	return nil
}
