package repositories

import (
	"microsrvc/merchant_svc/app/models"
	"microsrvc/merchant_svc/config/database"
)

func FindExistsMerchant(name string, userID uint) error {
	db := database.DB

	if err := db.Where("name = ?", name).
		Or("user_id = ?", userID).First(&models.Merchant{}).Error; err != nil {
		return err
	}

	return nil
}

func FetchMerchant() ([]*models.Merchant, error) {
	db := database.DB

	var merchants []*models.Merchant

	if err := db.Select([]string{"id", "name", "description", "city", "user_id"}).
		Find(&merchants).Error; err != nil {
		return nil, err
	}

	return merchants, nil
}

func GetMerchantById(id uint) (*models.Merchant, error) {
	db := database.DB
	var merchant *models.Merchant

	if err := db.Select([]string{"id", "name", "description", "city", "user_id"}).
		Where("id = ?", id).First(&merchant).Error; err != nil {
		return nil, err
	}

	return merchant, nil
}

func CreateMerchant(merchant *models.Merchant) error {
	db := database.DB

	if err := db.Create(merchant).Error; err != nil {
		return err
	}

	return nil
}
