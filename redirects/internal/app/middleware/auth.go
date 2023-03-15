package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

const roleAdmin = "admin"

func AdminCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := c.Next(); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}

		val := c.GetReqHeaders()["User-Role"]
		if strings.EqualFold(val, roleAdmin) {
			return nil
		}

		return c.Status(fiber.StatusForbidden).JSON("You are not admin")
	}
}
