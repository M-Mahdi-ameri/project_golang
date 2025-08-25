package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdi/library-api/internal/database"
	"github.com/mahdi/library-api/internal/handlers"
	"github.com/mahdi/library-api/internal/jobs"
	"github.com/mahdi/library-api/internal/middleware"
	"github.com/mahdi/library-api/internal/routs"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(middleware.Logger)

	api := app.Group("/api")
	routs.SetupBookRoutes(api)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello, library api!")
	})

	routs.SetupReportRoutes(api)

	app.Post("/signup", handlers.SingUp)
	app.Post("/login", handlers.Login)

	jobs.InitDispatcher(10)
	for i := 1; i <= 3; i++ {
		jobs.StartWorker(i)
	}

	jobs.InitReportDispatcher(5)
	for i := 1; i <= 2; i++ {
		jobs.StartReportWorker(i)
	}

	log.Fatal(app.Listen(":3000"))
}
