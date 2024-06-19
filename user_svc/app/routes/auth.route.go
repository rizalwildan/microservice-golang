package routes

import (
	"github.com/gofiber/fiber/v2"
	"microsrvc/user_svc/app/services"
	"microsrvc/user_svc/utils/middleware"
)

func AuthRoutes(app *fiber.App) {
	r := app.Group("/auth")

	r.Post("/signup", services.Signup)
	r.Post("/login", services.Login)
	r.Use(middleware.Auth).Get("/me", services.GetMe)
}
