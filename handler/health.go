package handler

import "github.com/gofiber/fiber/v2"

func (h *handler) Health(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
	})
}
