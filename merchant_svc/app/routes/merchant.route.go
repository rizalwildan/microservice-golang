package routes

import (
	"github.com/gofiber/fiber/v2"
	"microsrvc/merchant_svc/app/services"
)

func MerchantRoutes(app *fiber.App) {
	r := app.Group("/merchants")

	r.Post("/", services.CreateMerchant)
	r.Get("/", services.FetchMerchant)
	r.Get("/:id", services.GetMerchant)
	r.Get("/:merchantId/products", services.GetMerchantProduct)
}
