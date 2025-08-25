package routs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdi/library-api/internal/handlers"
	"github.com/mahdi/library-api/internal/middleware"
)

func SetupBookRoutes(router fiber.Router) {
	task := router.Group("/books", middleware.Protect())
	task.Post("/", handlers.CreateBook)
	task.Get("/:id", handlers.GetBook)
	task.Put("/:id", handlers.UpdateBook)
	task.Delete("/:id", handlers.DeleteBook)
}
