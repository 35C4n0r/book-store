package handler

import (
	"balkan/internal/models"
	"balkan/internal/models/validators"
	"balkan/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func AddBookHandler(c *gin.Context) {
	var req validators.BookAddUpdateRequestModel
	var err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var valid = service.ValidateISBN(req.ISBN)

	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_ISBN"})
		return
	}

	_, err = service.GetBookByISBN(req.ISBN)
	if err == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "BOOK_ALREADY_EXISTS"})
		return
	}

	var book = models.InventoryCRUD{
		ISBN:        req.ISBN,
		Title:       req.Title,
		Author:      req.Author,
		Description: req.Description,
		Content:     req.Content,
	}

	err = service.CreateBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERVAL_SERVER_ERROR"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "BOOK_CREATED_SUCCESSFULLY"})
}

func DeleteBookHandler(c *gin.Context) {
	var ISBN = c.Params.ByName("isbn")
	err := service.DeleteBook(ISBN)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERVAL_SERVER_ERROR"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "BOOK_DELETED_SUCCESSFULLY"})
}

func UpdateBookHandler(c *gin.Context) {
	var req validators.BookAddUpdateRequestModel
	var err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var valid = service.ValidateISBN(req.ISBN)

	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_ISBN"})
		return
	}

	_, err = service.GetBookByISBN(req.ISBN)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "BOOK_DOSENT_EXISTS"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERVAL_SERVER_ERROR"})
			return
		}
	}

	var book = models.InventoryCRUD{
		ISBN:        req.ISBN,
		Title:       req.Title,
		Author:      req.Author,
		Description: req.Description,
		Content:     req.Content,
	}

	err = service.UpdateBook(&book)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERVAL_SERVER_ERROR"})
		return
	}

	c.JSON(http.StatusOK, book)

}

func SearchBookHandler(c *gin.Context) {
	var isbn = c.DefaultQuery("isbn", "")
	var author = c.DefaultQuery("author", "")
	var title = c.DefaultQuery("title", "")
	dbBooks, err := service.SearchBooks(title, author, isbn)

	if err != nil {

	}

	var inventoryResponseItems []validators.InventoryResponse
	for _, crudItem := range dbBooks {
		respItem := validators.InventoryResponse{
			ISBN:        crudItem.ISBN,
			Title:       crudItem.Title,
			Author:      crudItem.Author,
			Description: crudItem.Description,
		}
		inventoryResponseItems = append(inventoryResponseItems, respItem)
	}
	c.JSON(http.StatusOK, validators.InventoryListResponse{Items: inventoryResponseItems})

}
