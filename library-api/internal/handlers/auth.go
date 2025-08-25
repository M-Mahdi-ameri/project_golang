package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdi/library-api/internal/database"
	"github.com/mahdi/library-api/internal/models"
	"github.com/mahdi/library-api/internal/utils"
)

var jwtSecret = []byte("my_secret_key")

func SingUp(c *fiber.Ctx) error {
	var data struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "invalid request",
			"detail": err.Error(),
		})
	}
	hashedPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not hash password",
		})
	}

	user := models.User{
		Username: data.Username,
		Email:    data.Email,
		Password: string(hashedPassword),
	}
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "could not create user",
		})
	}
	return c.JSON(fiber.Map{
		"message": "user created",
	})
}

func Login(c *fiber.Ctx) error {
	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	var user models.User

	if err := database.DB.Where("email = ?", data.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid credentials",
		})
	}

	if !utils.CheckPasswordHash(data.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid credentials",
		})
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not login",
		})
	}

	return c.JSON(fiber.Map{
		"token": token, "expires_in": time.Now().Add(time.Hour * 72)})

}
