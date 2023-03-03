package menu

import (
	"github.com/gofiber/fiber/v2"
	"goozt.org/breakstick/api/utils"
)

// Get all items in menu
func GetMenu(c *fiber.Ctx) error {
	menu := Menu{Items: []Item{}}

	// Get all items from store
	data, err := menuDB.GetAll()
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}

	// Decode binary map data to items
	for _, b := range data {
		var item Item
		Decode(b, &item)
		menu.Items = append(menu.Items, item)
	}

	return utils.SendJSON(c, menu)
}

// Delete all items in menu
func DeleteMenu(c *fiber.Ctx) error {
	err := menuDB.ResetStore()
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}

	return utils.SendJSON(c, fiber.Map{
		"deleted_menu": true,
	})
}
