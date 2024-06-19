package api

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"microsrvc/merchant_svc/app/types"
	"microsrvc/merchant_svc/config"
)

func CheckUser(id uint) (*types.User, error) {
	url := fmt.Sprintf("%s/users/%v", config.USER_SVC_URL, id)

	agent := fiber.Get(url)
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		log.Errorf("something happen with user_svc :: %v", errs)
		return nil, errs[0]
	}

	if statusCode == fiber.StatusNotFound {
		return nil, fiber.ErrNotFound
	}

	var response *types.UserApiResponse
	err := json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &types.User{
		ID:    response.Data.ID,
		Name:  response.Data.Name,
		Email: response.Data.Email,
	}, nil
}
