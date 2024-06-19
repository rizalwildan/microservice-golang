package main

import (
	"fmt"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"microsrvc/user_svc/app/models"
	"microsrvc/user_svc/app/routes"
	"microsrvc/user_svc/config"
	"microsrvc/user_svc/config/database"
)

func main() {
	// connect database
	database.Connect()
	database.Migrate(&models.User{})

	// swagger config
	cfg := swagger.Config{
		BasePath: "/users",
		FilePath: "./docs/swagger.yaml",
		Path:     "docs",
		Title:    "User API Docs",
		CacheAge: 1,
	}

	app := fiber.New()

	app.Use(logger.New())

	app.Use(swagger.New(cfg))

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "PONG!!!"})
	})

	routes.AuthRoutes(app)
	routes.UserRoutes(app)

	err := app.Listen(fmt.Sprintf(":%v", config.PORT))
	if err != nil {
		panic(err)
	}
}
