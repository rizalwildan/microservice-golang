package repositories

import (
	"microsrvc/product_svc/app/models"
	"microsrvc/product_svc/app/types"
	"microsrvc/product_svc/config/database"
)

func FetchProduct(query *types.Query) ([]*models.Product, error) {
	var products []*models.Product

	db := database.DB

	if query.MerchantId != 0 {
		db = db.Where("merchant_id = ?", query.MerchantId)
	}

	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func CreteProduct(product *models.Product) error {
	db := database.DB

	if err := db.Create(product).Error; err != nil {
		return err
	}

	return nil
}
