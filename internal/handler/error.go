package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	ctx.Status(code).JSON(fiber.Map{
		"Success": false,
		"Error":   err.Error(),
	})

	return nil
}
