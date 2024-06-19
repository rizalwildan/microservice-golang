package api

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"microsrvc/merchant_svc/app/types"
	"microsrvc/merchant_svc/config"
)

func GetProductByMerchantID(id uint) ([]*types.Product, error) {
	url := fmt.Sprintf("%s/products?merchant_id=%v", config.PRODUCT_SVC_URL, id)
	fmt.Println(url)

	agent := fiber.Get(url)
	agent.Add("X-Api-Key", config.PRODUCT_SVC_KEY)

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		log.Errorf("something happen with products_svc :: %v", errs)
		return nil, errs[0]
	}

	if statusCode == fiber.StatusNotFound {
		return nil, fiber.ErrNotFound
	}

	var response *types.ProductResponseAPI
	err := json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	var products []*types.Product
	for _, product := range response.Data {
		products = append(products, &types.Product{
			ID:       product.ID,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
		})
	}

	return products, nil
}
