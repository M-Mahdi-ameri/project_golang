package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mahdi/library-api/internal/utils"
)

var jwtSecret = []byte("my_secret_key")

func Protect() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorizaion")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing or malformed jwt",
			})
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := utils.ParseJWT(tokenStr)
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid or expired jwt",
			})
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Locals("userID", uint(claims["id"].(float64)))

		return c.Next()
	}
}
