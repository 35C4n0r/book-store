package service

import (
	"balkan/internal/database"
	"balkan/internal/models"
	"errors"
	"github.com/sirupsen/logrus"
)

func ValidateISBN(ISBN string) bool {
	//return IsValidISBN10(ISBN) || IsValidISBN13(ISBN)
	return true
}

func CreateBook(book models.InventoryCRUD) error {
	var result = database.DB.Create(&book)
	if result.Error != nil {
		return errors.New("BOOK_CREATION_FAIL")
	}
	return nil

}

func GetBookByISBN(ISBN string) (*models.InventoryCRUD, error) {
	var dbBook models.InventoryCRUD
	var result = database.DB.Where("isbn = ?", ISBN).First(&dbBook)
	if result.Error != nil {
		return nil, result.Error
	}
	return &dbBook, nil
}

func DeleteBook(ISBN string) error {
	var dbBook models.InventoryCRUD
	var result = database.DB.Debug().Where("isbn = ?", ISBN).Delete(&dbBook)
	if result.Error != nil {
		logrus.Errorf("Error while deleting book, %v", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		logrus.Errorf("The book with the requested ISBN dosen't exist")
		return errors.New("BOOK_NOT_EXISTS")
	}
	return nil
}

func UpdateBook(book *models.InventoryCRUD) error {
	var dbBook models.InventoryCRUD
	result := database.DB.Model(&dbBook).Where("isbn = ?", book.ISBN).Updates(book)
	if result.Error != nil {
		logrus.Errorf("Error while updating book, %v", result.Error)
		return result.Error
	}
	return nil
}

func SearchBooks(title string, author string, isbn string) ([]models.InventoryCRUD, error) {
	var books []models.InventoryCRUD
	query := database.DB.Model(&models.InventoryCRUD{})

	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}

	if author != "" {
		query = query.Where("author LIKE ?", "%"+author+"%")
	}

	if isbn != "" {
		query = query.Where("isbn LIKE ?", "%"+isbn+"%")
	}

	err := query.Find(&books).Error
	if err != nil {
		logrus.Errorf("Error While Searching Books, %v", err)
		return nil, err
	}

	return books, nil
}
