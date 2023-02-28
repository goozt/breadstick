package main

import (
	"github.com/gofiber/fiber/v2"
	"goozt.org/breakstick/api"
	"goozt.org/breakstick/api/menu"
	"goozt.org/breakstick/api/utils"
)

func RegisterHandlers(a *api.API) {
	menu.HandleRoutes(a.Router)
	HealthHandler(a.App)
}

func HealthHandler(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return utils.SendJSON(c, fiber.Map{
			"status": "ok",
		})
	})
}
