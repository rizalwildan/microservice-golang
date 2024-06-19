package routes

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"microsrvc/product_svc/app/services"
	"microsrvc/product_svc/utils/middleware"
)

func ProductRoutes(app *fiber.App) {
	r := app.Group("/products").Use(keyauth.New(keyauth.Config{
		KeyLookup: "header:X-Api-Key",
		Validator: middleware.ValidateAPIKey,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if errors.Is(err, keyauth.ErrMissingOrMalformedAPIKey) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": err.Error()})
			}

			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": "Invalid or expired API Key"})
		},
	}))

	r.Post("/", services.CreateProduct)
	r.Get("/", services.FetchProduct)
}
