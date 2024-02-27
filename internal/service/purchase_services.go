package service

import (
	"balkan/internal/database"
	"balkan/internal/models"
	"github.com/sirupsen/logrus"
)

func MakePurchase(purchase models.PurchaseCRUD) error {
	var result = database.DB.Create(&purchase)
	if result.Error != nil {
		logrus.Errorf("Error while creating a new user, %v", result.Error)
		return result.Error
	}
	return nil
}

func HasUserPurchasedBook(userID, ISBN string) (bool, error) {
	var count int64
	err := database.DB.Model(&models.PurchaseCRUD{}).Where("user_id = ? AND isbn = ?", userID, ISBN).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func GetBookContentByISBN(ISBN string) ([]byte, error) {
	var book models.InventoryCRUD
	err := database.DB.Where("isbn = ?", ISBN).First(&book).Error
	if err != nil {
		return nil, err
	}
	return book.Content, nil
}

func GetPurchaseByUserIdAndISBN(ISBN string, UserId string) (*models.PurchaseCRUD, error) {
	var dbPurchase models.PurchaseCRUD
	result := database.DB.Where("user_id = ?", UserId).Where("isbn = ?", ISBN).First(&dbPurchase)
	if result.Error != nil {
		return nil, result.Error
	}
	return &dbPurchase, nil
}
