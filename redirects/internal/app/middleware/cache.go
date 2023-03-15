package middleware

import (
	"github.com/alikhanturusbekov/redirects/internal/app/cache"
	"github.com/gofiber/fiber/v2"
)

func CacheCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		link := c.Query("link")
		if link == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Provide link parameter"})
		}

		val, cached := cache.Cr.Get(link)
		if cached {
			return c.Status(fiber.StatusMovedPermanently).JSON(fiber.Map{"status": "success", "link": val})
		}

		if err := c.Next(); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}

		return nil
	}
}
