package services

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"microsrvc/user_svc/app/models"
	"microsrvc/user_svc/app/repositories"
	"microsrvc/user_svc/app/types"
	"microsrvc/user_svc/utils"
	"microsrvc/user_svc/utils/jwt"
	"microsrvc/user_svc/utils/password"
)

func Login(ctx *fiber.Ctx) error {
	body := new(types.LoginDTO)

	if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "validation failed",
			"error":   err.Error(),
		})
	}

	user, err := repositories.FindUserByEmail(body.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "invalid email or password",
		})
	}

	if err := password.Verify(user.Password, body.Password); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "invalid email or password",
		})
	}

	t := jwt.GenerateToken(&jwt.TokenPayload{
		ID: user.ID,
	})

	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    fiber.Map{"token": t},
	})
}

func Signup(ctx *fiber.Ctx) error {
	body := new(types.SignupDTO)

	if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "validation failed",
			"errors":  err.Error(),
		})
	}

	_, err := repositories.FindUserByEmail(body.Email)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"status": false, "message": "email already exists"})
	}

	user := &models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: password.Generate(body.Password),
	}

	err = repositories.CreateUser(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": false, "message": "failed to create user"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "created",
	})
}

func GetMe(ctx *fiber.Ctx) error {
	user, err := repositories.FindUserById(utils.GetUser(ctx))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data": &types.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	})
}
