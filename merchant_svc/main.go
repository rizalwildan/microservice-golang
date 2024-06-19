package main

import (
	"fmt"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"microsrvc/merchant_svc/app/models"
	"microsrvc/merchant_svc/app/routes"
	"microsrvc/merchant_svc/config"
	"microsrvc/merchant_svc/config/database"
)

func main() {
	// connect database
	database.Connect()
	database.Migrate(&models.Merchant{})

	// swagger config
	cfg := swagger.Config{
		BasePath: "/merchants",
		FilePath: "./docs/swagger.yaml",
		Path:     "docs",
		Title:    "Products API Docs",
		CacheAge: 1,
	}

	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover2.New())
	app.Use(swagger.New(cfg))

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("PONG from merchant_svc")
	})

	app.Get("/panic", func(c *fiber.Ctx) error {
		panic("PANIC from merchant_svc")
	})

	routes.MerchantRoutes(app)

	err := app.Listen(fmt.Sprintf(":%v", config.PORT))
	if err != nil {
		panic(err)
	}
}
