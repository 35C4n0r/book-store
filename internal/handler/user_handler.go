package handler

import (
	"balkan/internal/models/validators"
	"balkan/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeactivateUserAccountHandler(c *gin.Context) {
	var userId, _ = c.Get("UserId")
	actualUserID, _ := userId.(string)

	var err = service.DeactivateUser(actualUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL_SERVER_ERROR"})
		return
	}
	res := validators.DeactivationResponseModel{Message: "USER_DEACTIVATED"}
	c.JSON(http.StatusOK, res)
}

func DeleteUserAccountHandler(c *gin.Context) {
	var userId, _ = c.Get("UserId")
	actualUserID, _ := userId.(string)
	var err = service.DeleteUser(actualUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL_SERVER_ERROR"})
		return
	}
	res := validators.DeleteResponseModel{Message: "USER_DELETED"}
	c.JSON(http.StatusOK, res)
}
