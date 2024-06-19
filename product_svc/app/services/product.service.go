package services

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"microsrvc/product_svc/api"
	"microsrvc/product_svc/app/models"
	"microsrvc/product_svc/app/repositories"
	"microsrvc/product_svc/app/types"
	"microsrvc/product_svc/utils"
)

func FetchProduct(ctx *fiber.Ctx) error {
	query := new(types.Query)

	if err := ctx.QueryParser(query); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	products, err := repositories.FetchProduct(query)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	var response []*types.ProductResponse

	for _, product := range products {
		response = append(response, &types.ProductResponse{
			ID:         product.ID,
			Name:       product.Name,
			Price:      product.Price,
			Quantity:   product.Quantity,
			MerchantID: product.MerchantID,
		})
	}

	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    response,
	})
}

func CreateProduct(ctx *fiber.Ctx) error {
	body := new(types.ProductDTO)

	if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "validation failed",
			"error":   err.Error(),
		})
	}

	_, err := api.CheckMerchant(body.MerchantID)
	if err != nil {
		if errors.Is(err, fiber.ErrNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "merchant not found",
			})
		}

		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "something happen with merchant service",
		})
	}

	product := &models.Product{
		Name:        body.Name,
		Description: body.Description,
		Quantity:    body.Quantity,
		Price:       body.Price,
		MerchantID:  body.MerchantID,
	}

	err = repositories.CreteProduct(product)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "error creating merchant",
		})
	}

	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "created",
	})
}
