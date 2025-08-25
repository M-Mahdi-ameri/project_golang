package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdi/library-api/internal/database"
	"github.com/mahdi/library-api/internal/jobs"
	"github.com/mahdi/library-api/internal/models"
)

func CreateBook(c *fiber.Ctx) error {
	var book models.Book

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	userID := c.Locals("userID").(uint)
	book.UserID = userID

	if err := database.DB.Create(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not create book",
		})
	}

	newJob := jobs.Job{
		ID:      int(book.ID),
		Payload: book.Title,
	}
	jobs.JobQueue <- newJob

	return c.Status(201).JSON(book)
}

func GetBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid book id ",
		})
	}
	userID := c.Locals("userID").(uint)
	var book models.Book

	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&book).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "book not found",
		})
	}
	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid book id ",
		})
	}
	userID := c.Locals("userID").(uint)

	var book models.Book

	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&book).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "book not found",
		})
	}

	var data models.Book

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	book.Title = data.Title
	book.Author = data.Author
	book.Descripion = data.Descripion

	if err := database.DB.Save(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not update book",
		})
	}

	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid book id ",
		})
	}
	userID := c.Locals("userID").(uint)
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Book{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not delete book",
		})
	}

	return c.JSON(fiber.Map{
		"message": "book deleted",
	})
}
