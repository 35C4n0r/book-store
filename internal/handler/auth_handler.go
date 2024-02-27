package handler

import (
	"balkan/internal/auth"
	"balkan/internal/models"
	"balkan/internal/models/validators"
	"balkan/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
)

func LoginHandler(c *gin.Context) {

	var req validators.LoginRequestModel
	var err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var dbUser, DatabaseErr = service.VerifyUser(req.Email, req.Password)
	if DatabaseErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "INVALID_USERNAME/PASSWORD"})
		return
	}
	var JWT, JWTErr = auth.GenerateJWT(dbUser)
	if JWTErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL_SERVER_ERROR"})
		return
	}

	var ResUser = validators.User{
		Username: dbUser.Username,
		Email:    dbUser.Email,
	}

	var response = validators.LoginResponseModel{
		Token:     JWT,
		User:      ResUser,
		ExpiresIn: 3600,
	}

	c.JSON(http.StatusOK, response)
}

func RegisterHandler(c *gin.Context) {

	var req validators.RegisterRequestModel
	var err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Checking if account already exists
	var _, userExistsError = service.GetUserByEmail(req.Email)
	if userExistsError == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "USER_ALREADY_EXISTS"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL_SERVER_ERROR"})
		return
	}
	newUUID := uuid.New()
	pgUUID := pgtype.UUID{
		Bytes: newUUID,
		Valid: true,
	}
	var user = models.UserCRUD{
		ID:           pgUUID,
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}
	CRUDErr := service.CreateUser(user)
	if CRUDErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERVAL_SERVER_ERROR"})
		return
	}

	var ResUser = validators.User{
		Email:    req.Email,
		Username: req.Username,
	}

	var response = validators.RegisterResponseModel{
		Message: "USER_REGISTERED_SUCCESSFULLY",
		User:    ResUser,
	}

	c.JSON(http.StatusCreated, response)
}

func RegisterAdminHandler(c *gin.Context) {
	var req validators.AdminRequest
	var err = c.ShouldBindJSON(&req)
	if req.AdminCode != os.Getenv("ADMIN_SECRET") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "UNATHORIZED"})
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	UserID, _ := c.Get("UserId")
	StringUserID, _ := UserID.(string)
	var UUIDUserId pgtype.UUID
	err = UUIDUserId.Scan(StringUserID)
	err = service.RegisterAdmin(StringUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL_SERVER_ERROR"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ADMIN_CREATED_SUCCESSFULLY"})
}
