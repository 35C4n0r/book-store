package service

import (
	"balkan/internal/database"
	"balkan/internal/models"
)

func AddReview(dbReview models.ReviewCRUD) error {
	err := database.DB.Create(&dbReview).Error
	if err != nil {
		return err
	}
	return nil
}

func GetReviewByUserID(userID string) ([]models.ReviewCRUD, error) {
	var cartItems []models.ReviewCRUD
	if err := database.DB.Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		return nil, err
	}
	return cartItems, nil
}
