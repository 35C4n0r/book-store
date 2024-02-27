package auth

import (
	"balkan/internal/models"
	"balkan/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"

	//"github.com/golang-jwt/jwt/v5"
	"time"
)

var SecretKey = []byte(os.Getenv("SERVER_SECRET"))

func GenerateJWT(user *models.UserCRUD) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(60 * time.Minute).Unix()
	claims["authorized"] = true
	claims["Email"] = user.Email
	claims["Username"] = user.Username
	claims["UserId"] = user.ID
	claims["UserRegisteredAt"] = user.CreatedAt

	// Sign and get the complete encoded token as a string using the secret

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		logrus.Errorf("Failed to generate JWT, error: %v", err)
		return "", err
	}
	return tokenString, nil
}

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			logrus.Errorf("missing authourization token in headers.")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header provided"})
			c.Abort()
			return
		}

		tokenString := authHeader[len(BearerSchema):]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				logrus.Errorf("unexpected signing method: %v", token.Header["alg"])
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return SecretKey, nil
		})

		if err != nil {
			logrus.Errorf("failed to parse/validate JWT token: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "INVALID_TOKEN"})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		var userId = claims["UserId"]
		dbUser, err := service.GetUserByID(claims["UserId"].(string))
		var isAdmin = dbUser.IsAdmin
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL_SERVER_ERROR"})
			c.Abort()
			return
		}
		if !dbUser.IsActive {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "ACCOUNT_DEACTIVE"})
			c.Abort()
			return
		}
		if ok && token.Valid {
			c.Set("UserId", userId)
			c.Set("isAdmin", isAdmin)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "INVALID_TOKEN"})
			c.Abort()
			return
		}
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var isAdmin, _ = c.Get("isAdmin")
		isAdminActual, _ := isAdmin.(bool)
		if !isAdminActual {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "ACCESS_NOT_ALLOWED"})
			c.Abort()
			return
		}
		c.Next()
	}
}
