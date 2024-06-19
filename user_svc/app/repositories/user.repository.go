package repositories

import (
	"microsrvc/user_svc/app/models"
	"microsrvc/user_svc/config/database"
)

func FindUserByEmail(email string) (*models.User, error) {
	db := database.DB

	var user models.User

	if err := db.Where(&models.User{Email: email}).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func FindUserById(id uint) (*models.User, error) {
	db := database.DB

	var user models.User

	if err := db.Select([]string{"id", "name", "email"}).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(user *models.User) error {
	db := database.DB

	err := db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}
