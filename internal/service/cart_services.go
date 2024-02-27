package service

import (
	"balkan/internal/database"
	"balkan/internal/models"
	"errors"
	"github.com/sirupsen/logrus"
)

func AddToCart(cartItem models.CartCRUD) error {
	err := database.DB.Create(&cartItem).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCartItemByUserAndISBN(UserId string, ISBN string) (*models.CartCRUD, error) {
	var cartItem models.CartCRUD
	result := database.DB.Where("user_id = ?", UserId).Where("isbn = ?", ISBN).First(&cartItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &cartItem, nil
}

func UpdateCart(cartItem models.CartCRUD) error {
	var dbCart models.CartCRUD
	result := database.DB.Model(&dbCart).Where("user_id = ?", cartItem.UserID).Where("isbn = ?", cartItem.ISBN).Update("quantity", cartItem.Quantity)
	if result.Error != nil {
		logrus.Errorf("Error while updating cart, %v", result.Error)
		return result.Error
	}
	return nil
}

func DeleteCart(ISBN string, UserId string) error {
	var dbCart models.CartCRUD
	var result = database.DB.Debug().Where("isbn = ?", ISBN).Where("user_id = ?", UserId).Delete(&dbCart)
	if result.Error != nil {
		logrus.Errorf("Error while deleting cart item, %v", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		logrus.Errorf("The cartItem with the requested ISBN and UserId dosen't exist")
		return errors.New("CART_ITEM_NOT_EXISTS")
	}
	return nil
}

func GetCartItemsByUserID(userID string) ([]models.CartCRUD, error) {
	var cartItems []models.CartCRUD
	if err := database.DB.Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		return nil, err
	}
	return cartItems, nil
}
