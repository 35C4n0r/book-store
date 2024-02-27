package handler

import (
	"balkan/internal/models"
	"balkan/internal/models/validators"
	"balkan/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
)

func PurchaseBookHandler(c *gin.Context) {
	var req validators.PurchaseRequest
	var err = c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	UserID, _ := c.Get("UserId")
	StringUserID, _ := UserID.(string)
	var UUIDUserId pgtype.UUID
	err = UUIDUserId.Scan(StringUserID)
	if err != nil {
		return
	}

	_, err = service.GetPurchaseByUserIdAndISBN(req.ISBN, StringUserID)
	if err == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ALREADY_PURCHASED"})
		return
	}

	newUUID := uuid.New()
	pgUUID := pgtype.UUID{
		Bytes: newUUID,
		Valid: true,
	}
	var purchase = models.PurchaseCRUD{
		ID:     pgUUID,
		UserID: UUIDUserId,
		ISBN:   req.ISBN,
	}
	CRUDErr := service.MakePurchase(purchase)
	if CRUDErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERVAL_SERVER_ERROR"})
		return
	}

	var response = validators.PurchaseResponse{
		Message: "BOOK_PURCHASED_SUCCESSFULLY",
		ISBN:    req.ISBN,
	}

	c.JSON(http.StatusOK, response)
}

func DownloadBookHandler(c *gin.Context) {

	ISBN := c.Params.ByName("isbn")
	UserID, _ := c.Get("UserId")
	StringUserID, _ := UserID.(string)
	var UUIDUserId pgtype.UUID
	err := UUIDUserId.Scan(StringUserID)
	if err != nil {
		return
	}

	// Check if the user has purchased the book
	hasPurchased, err := service.HasUserPurchasedBook(StringUserID, ISBN)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL_SERVER_ERROR"})
		return
	}
	if !hasPurchased {
		c.JSON(http.StatusForbidden, gin.H{"error": "BOOK_NOT_PURCHASED"})
		return
	}

	// Fetch book content
	content, err := service.GetBookContentByISBN(ISBN)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL_SERVER_ERROR"})
		return
	}

	// Return the content as a file download
	c.Data(http.StatusOK, "application/octet-stream", content)
}
