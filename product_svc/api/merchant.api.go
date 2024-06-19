package api

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"microsrvc/product_svc/app/types"
	"microsrvc/product_svc/config"
)

func CheckMerchant(id uint) (*types.Merchant, error) {
	url := fmt.Sprintf("%s/merchants/%v", config.MERCHANT_SVC_URL, id)

	agent := fiber.Get(url)
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		log.Errorf("something happen with user_svc :: %v", errs)
		return nil, errs[0]
	}

	if statusCode == fiber.StatusNotFound {
		return nil, fiber.ErrNotFound
	}

	var response *types.MerchantApiResponse
	err := json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &types.Merchant{
		ID:   response.Data.ID,
		Name: response.Data.Name,
		City: response.Data.City,
	}, nil
}
