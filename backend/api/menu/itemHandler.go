package menu

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"goozt.org/breakstick/api/utils"
)

// Create a new item
func CreateItem(c *fiber.Ctx) error {
	reqItem := new(Item)

	// Parse the request parameters
	if err := c.BodyParser(reqItem); err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}

	// Copy data with request data
	item := NewItem()
	item.Copy(reqItem)

	// Validate the item
	errorList := ValidateItem(c, *item)
	if len(errorList) > 0 {
		return utils.SendValidationError(c, fiber.StatusBadRequest, errorList)
	}

	// Check if item exists
	if itemExists := menuDB.ValueExists(item.Name, func(data []byte) (string, error) {
		var decodedItem Item
		err := Decode(data, &decodedItem)
		return decodedItem.Name, err
	}); itemExists {
		return utils.SendError(c,
			fiber.StatusNotAcceptable,
			errors.New("item already exist"),
		)
	}

	// Encode to bytes
	encoded, err := item.Encode()
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}

	// Add the item to the store
	err = menuDB.Set(item.ID, encoded)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}

	return utils.SendJSON(c, item)
}

// Get the item with the given ID
func GetItem(c *fiber.Ctx) error {

	// Get ID from URL
	id := c.Params("id")
	if id == "" {
		return utils.SendError(c, fiber.StatusBadRequest)
	}

	// Get item from store
	data, err := menuDB.Get(id)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}
	if data == nil {
		return utils.SendError(c, fiber.StatusNotFound, errors.New("item not found"))
	}

	// Decode binary data to item
	var item Item
	err = Decode(data, &item)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}

	return utils.SendJSON(c, item)
}

// Get the items with the given ID
func UpdateItem(c *fiber.Ctx) error {
	// Get ID from URL
	id := c.Params("id")
	if id == "" {
		return utils.SendError(c, fiber.StatusBadRequest)
	}

	// Parse the request parameters
	reqItem := new(Item)
	if err := c.BodyParser(reqItem); err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}

	// Verify the item exists
	if !menuDB.KeyExists(id) {
		return utils.SendError(c,
			fiber.StatusBadRequest,
			errors.New("Invalid ID or Item does not exists"),
		)
	}

	// Get item from store
	data, err := menuDB.Get(id)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}
	if data == nil {
		return utils.SendError(c, fiber.StatusNotFound, errors.New("item not found"))
	}

	// Decode binary data to item
	var item Item
	err = Decode(data, &item)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}

	// Copy data with request data
	item.Copy(reqItem)

	// Validate the item
	errorList := ValidateItem(c, item)
	if len(errorList) > 0 {
		return utils.SendValidationError(c, fiber.StatusBadRequest, errorList)
	}

	// Encode to bytes
	encodedItem, err := item.Encode()
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}

	// Add the item to the store
	err = menuDB.Set(id, encodedItem)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}

	return utils.SendJSON(c, item)
}

// Delete the item with given ID
func DeleteItem(c *fiber.Ctx) error {
	// Get ID from URL
	id := c.Params("id")
	if id == "" {
		return utils.SendError(c, fiber.StatusBadRequest)
	}

	// Get item from store
	if menuDB.KeyExists(id) {
		err := menuDB.Delete(id)
		if err != nil {
			return utils.SendError(c, fiber.StatusInternalServerError, err)
		} else {
			return utils.SendJSON(c, fiber.Map{"deleted_item": id})
		}
	}

	return utils.SendError(c, fiber.StatusBadRequest)
}
