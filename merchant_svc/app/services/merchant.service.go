package services

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
	"microsrvc/merchant_svc/api"
	"microsrvc/merchant_svc/app/models"
	"microsrvc/merchant_svc/app/repositories"
	"microsrvc/merchant_svc/app/types"
	"microsrvc/merchant_svc/utils"
	"strconv"
)

func FetchMerchant(ctx *fiber.Ctx) error {
	merchants, err := repositories.FetchMerchant()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "something went wrong",
			"errors":  err,
		})
	}

	var response []*types.MerchantResponse
	for _, merchant := range merchants {
		response = append(response, &types.MerchantResponse{
			ID:          merchant.ID,
			Name:        merchant.Name,
			Description: merchant.Description,
			City:        merchant.City,
			UserID:      merchant.UserID,
		})
	}

	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    response,
	})
}

func GetMerchant(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	merchant, err := repositories.GetMerchantById(uint(id))
	if err != nil {
		if errors.Is(err, fiber.ErrNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "merchant not found",
			})
		}

		return fiber.ErrInternalServerError
	}

	user, err := api.CheckUser(uint(id))
	if err != nil {
		if !errors.Is(err, fiber.ErrNotFound) {
			log.Errorf("something happend with user_svc :: %v", err)
		}
	}

	merchantResponse := &types.MerchantResponse{
		ID:          merchant.ID,
		Name:        merchant.Name,
		Description: merchant.Description,
		City:        merchant.City,
		UserID:      merchant.UserID,
		Owner:       user,
	}

	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "oke",
		"data":    merchantResponse,
	})
}

func CreateMerchant(ctx *fiber.Ctx) error {
	body := new(types.MerchantDTO)

	if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "validation failed",
			"error":   err.Error(),
		})
	}

	err := repositories.FindExistsMerchant(body.Name, body.UserID)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"success": false,
			"message": "merchant already exists",
		})
	}

	_, err = api.CheckUser(body.UserID)
	if err != nil {
		if errors.Is(err, fiber.ErrNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "user not found",
			})
		}

		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "something happen with user service",
		})
	}

	merchant := &models.Merchant{
		Name:        body.Name,
		Description: body.Description,
		City:        body.City,
		UserID:      body.UserID,
	}

	err = repositories.CreateMerchant(merchant)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "error creating merchant",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "created",
	})
}

func GetMerchantProduct(ctx *fiber.Ctx) error {
	merchantId, _ := strconv.Atoi(ctx.Params("merchantId"))

	merchant, err := repositories.GetMerchantById(uint(merchantId))
	if err != nil {
		if errors.Is(err, fiber.ErrNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "merchant not found",
			})
		}

		return fiber.ErrInternalServerError
	}

	user, err := api.CheckUser(uint(merchantId))
	if err != nil {
		if !errors.Is(err, fiber.ErrNotFound) {
			log.Errorf("something happend with user_svc :: %v", err)
		}
	}

	products, err := api.GetProductByMerchantID(uint(merchantId))
	if err != nil {
		if !errors.Is(err, fiber.ErrNotFound) {
			log.Errorf("something happend with product_svc :: %v", err)
		}
	}

	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data": &types.MerchantResponse{
			ID:          merchant.ID,
			Name:        merchant.Name,
			Description: merchant.Description,
			City:        merchant.City,
			Owner:       user,
			Products:    products,
		},
	})
}
