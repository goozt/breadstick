package menu

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// Validate Item struct
func ValidateItem(c *fiber.Ctx, item Item) []fiber.Map {
	var errorMap []fiber.Map
	err := validate.Struct(item)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			m := fiber.Map{
				"failed_field": err.StructNamespace(),
				"tag":          err.Tag(),
				"value":        err.Param(),
			}
			errorMap = append(errorMap, m)
		}
	}

	return errorMap
}
