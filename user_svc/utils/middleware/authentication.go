package middleware

import (
	"github.com/gofiber/fiber/v2"
	"microsrvc/user_svc/utils/jwt"
	"strings"
)

func Auth(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": false, "message": "Unauthorized"})
	}

	// split header
	chunks := strings.Split(token, " ")

	if len(chunks) < 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": false, "message": "Unauthorized"})
	}

	user, err := jwt.VerifyToken(chunks[1])
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": false, "message": "Unauthorized"})
	}

	c.Locals("USER", user.ID)

	return c.Next()
}
