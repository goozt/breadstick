package menu

import "github.com/gofiber/fiber/v2"

// Load menu routes
func HandleRoutes(r fiber.Router) {
	route := r.Group("/menu")
	route.Get("/", GetMenu)
	route.Get("/:id", GetItem)
	route.Post("/", CreateItem)
	route.Put("/:id", UpdateItem)
	route.Delete("/", DeleteMenu)
	route.Delete("/:id", DeleteItem)
}
