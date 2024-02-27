package handler

import (
	"balkan/internal/models"
	"balkan/internal/models/validators"
	"balkan/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"gorm.io/gorm"
	"net/http"
)

func AddToCartHandler(c *gin.Context) {
	var req validators.CartAddUpdateRequestModel
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

	_, err = service.GetCartItemByUserAndISBN(StringUserID, req.ISBN)
	if err == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ITEM_ALREADY_IN_CART"})
		return
	}

	var cartItem = models.CartCRUD{
		UserID:   UUIDUserId,
		ISBN:     req.ISBN,
		Quantity: req.Quantity,
	}

	err = service.AddToCart(cartItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL_SERVER_ERROR"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ITEM_ADDED_SUCCESSFULLY"})
}

func UpdateCartItemHandler(c *gin.Context) {
	var req validators.CartAddUpdateRequestModel
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

	_, err = service.GetCartItemByUserAndISBN(StringUserID, req.ISBN)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "CART_ITEM_DOSENT_EXIST"})
			return
		} else {
		}
	}

	var cartItem = models.CartCRUD{
		UserID:   UUIDUserId,
		ISBN:     req.ISBN,
		Quantity: req.Quantity,
	}

	err = service.UpdateCart(cartItem)
	if err != nil {

	}

	response := validators.CartResponse{
		ISBN:     req.ISBN,
		Quantity: req.Quantity,
	}

	c.JSON(http.StatusOK, response)
}

func DeleteCartHandler(c *gin.Context) {
	var ISBN = c.Params.ByName("isbn")
	UserID, _ := c.Get("UserId")
	StringUserID, _ := UserID.(string)
	var UUIDUserId pgtype.UUID
	var err = UUIDUserId.Scan(StringUserID)
	if err != nil {
		return
	}

	_, err = service.GetCartItemByUserAndISBN(StringUserID, ISBN)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "CART_ITEM_DOSENT_EXIST"})
		return
	}

	err = service.DeleteCart(ISBN, StringUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERVAL_SERVER_ERROR"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "CART_DELETED_SUCCESSFULLY"})
}

func GetCartHandler(c *gin.Context) {
	userIdInterface, exists := c.Get("UserId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userId, ok := userIdInterface.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	cartItems, _ := service.GetCartItemsByUserID(userId)

	var responseItems []validators.CartItemResponse
	for _, item := range cartItems {
		responseItem := validators.CartItemResponse{
			ID:       item.ID,
			ISBN:     item.ISBN,
			Quantity: item.Quantity,
		}
		responseItems = append(responseItems, responseItem)
	}

	c.JSON(http.StatusOK, gin.H{"cart_items": responseItems})
}
