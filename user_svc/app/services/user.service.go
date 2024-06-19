package services

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"microsrvc/user_svc/app/repositories"
	"strconv"
)

func GetUser(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	user, err := repositories.FindUserById(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "user not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    user,
	})
}
