package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/alikhanturusbekov/redirects/internal/app/handler"
	"github.com/alikhanturusbekov/redirects/internal/app/middleware"
)

func setupRoutes(app *fiber.App) {
	app.Use(logger.New())

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Redirects Project",
		})
	})

	redirects := app.Group("/redirects")
	redirects.Use(middleware.CacheCheck())
	redirects.Get("/", handler.Redirect)

	admin := app.Group("/admin")
	admin.Use(middleware.AdminCheck())

	admin.Get("/redirects", handler.GetRedirects)
	admin.Get("/redirects/:redirectId", handler.GetRedirectById)
	admin.Post("/redirects", handler.CreateRedirect)
	admin.Patch("/redirects/:redirectId", handler.UpdateRedirect)
	admin.Delete("/redirects/:redirectId", handler.DeleteRedirect)
}
