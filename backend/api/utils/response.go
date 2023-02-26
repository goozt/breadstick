package utils

import (
	"github.com/gofiber/fiber/v2"
)

// JSON response
func SendJSON[T any](c *fiber.Ctx, data T) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    data,
	})
}

// Error response
func SendError(c *fiber.Ctx, status int, errorData ...error) error {
	var errs []string
	for _, err := range errorData {
		errs = append(errs, err.Error())
	}

	return c.Status(status).JSON(fiber.Map{
		"message": "error",
		"error":   fiber.NewError(status),
		"details": errs,
	})
}

// Validation Error response
func SendValidationError(c *fiber.Ctx, status int, errorData ...[]fiber.Map) error {
	var errDetails []fiber.Map

	if len(errorData) > 0 {
		errDetails = errorData[0]
	}

	return c.Status(status).JSON(fiber.Map{
		"message": "error",
		"error":   fiber.NewError(status),
		"details": errDetails,
	})
}
