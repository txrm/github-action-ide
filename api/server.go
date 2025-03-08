package api

// remove later

import (
	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	app := fiber.New()

	app.Get("/inspect/:action", func(c *fiber.Ctx) error {
		action := c.Params("action")
		return c.JSON(fiber.Map{
			"action":  action,
			"details": "This is where the GitHub Action details would be fetched.",
		})
	})

	app.Listen(":8080")
}
