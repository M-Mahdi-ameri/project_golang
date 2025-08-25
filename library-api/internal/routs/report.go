package routs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdi/library-api/internal/jobs"
)

var reportCounter = 0

func SetupReportRoutes(router fiber.Router) {
	router.Get("/report", func(c *fiber.Ctx) error {
		reportCounter++
		newJob := jobs.ReportJob{RequestID: reportCounter}
		jobs.ReportQueue <- newJob

		return c.JSON(fiber.Map{
			"status":  "queued",
			"message": "report generation started in bsckgrand",
			"job_id":  reportCounter,
		})
	})
}
