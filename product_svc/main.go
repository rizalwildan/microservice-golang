package main

import (
	"fmt"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"microsrvc/product_svc/app/models"
	"microsrvc/product_svc/app/routes"
	"microsrvc/product_svc/config"
	"microsrvc/product_svc/config/database"
)

func main() {
	// connect database
	database.Connect()
	database.Migrate(&models.Product{})

	// swagger config
	cfg := swagger.Config{
		BasePath: "/products",
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
		return c.SendString("PONG from product_svc")
	})

	routes.ProductRoutes(app)

	err := app.Listen(fmt.Sprintf(":%s", config.PORT))
	if err != nil {
		panic(err)
	}
}
