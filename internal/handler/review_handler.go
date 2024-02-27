package handler

import (
	"balkan/internal/models"
	"balkan/internal/models/validators"
	"balkan/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
)

func AddReviewHandler(c *gin.Context) {
	var req validators.ReviewAddRequest
	var err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	UserID, _ := c.Get("UserId")
	StringUserID, _ := UserID.(string)
	var UUIDUserId pgtype.UUID
	err = UUIDUserId.Scan(StringUserID)
	dbReview := models.ReviewCRUD{
		UserID: UUIDUserId,
		ISBN:   req.ISBN,
		Review: req.Review,
	}
	err = service.AddReview(dbReview)
	if err != nil {

	}
	c.JSON(http.StatusOK, gin.H{"message": "REVIEW_ADDED_SUCCESSFULLY"})
}

func GetReviewHandler(c *gin.Context) {
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

	cartItems, _ := service.GetReviewByUserID(userId)

	var responseItems []validators.ReviewResponse
	for _, item := range cartItems {
		responseItem := validators.ReviewResponse{
			ISBN:   item.ISBN,
			Review: item.Review,
		}
		responseItems = append(responseItems, responseItem)
	}

	c.JSON(http.StatusOK, gin.H{"cart_items": responseItems})
}
