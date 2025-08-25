package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Logger(c *fiber.Ctx) error {
	fmt.Println("request recived on path: ", c.Path())
	return c.Next()
}
