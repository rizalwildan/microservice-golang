package routes

import (
	"github.com/gofiber/fiber/v2"
	"microsrvc/user_svc/app/services"
)

func UserRoutes(app *fiber.App) {
	r := app.Group("/users")

	r.Get("/:id", services.GetUser)
}
