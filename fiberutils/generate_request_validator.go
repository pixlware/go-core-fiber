package fiberutils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type ValidationError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Value   any    `json:"value"`
	Message string `json:"message"`
}

const REQ_BODY_KEY = "reqBody"

func GenerateRequestValidator[T any](errorCode string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body T

		if err := c.BodyParser(&body); err != nil {
			responseBody := NewErrorResponseBody(fiber.StatusBadRequest, "Invalid request body", err, errorCode)
			return c.Status(fiber.StatusBadRequest).JSON(responseBody)
		}

		if err := validate.Struct(body); err != nil {
			var validationErrors []ValidationError
			if errs, ok := err.(validator.ValidationErrors); ok {
				for _, ve := range errs {
					validationErrors = append(validationErrors, ValidationError{
						Field:   ve.Field(),
						Tag:     ve.Tag(),
						Value:   ve.Value(),
						Message: ve.Field() + " is " + ve.Tag(),
					})
				}
			} else {
				// Some other kind of error
				responseBody := NewErrorResponseBody(fiber.StatusBadRequest, "Validation error", err, errorCode)
				return c.Status(fiber.StatusBadRequest).JSON(responseBody)
			}

			// Return array of validation errors
			responseBody := NewErrorResponseBody(fiber.StatusBadRequest, "Validation failed", validationErrors, errorCode)
			return c.Status(fiber.StatusBadRequest).JSON(responseBody)
		}

		// Store validated struct
		c.Locals(REQ_BODY_KEY, body)
		return c.Next()
	}
}
