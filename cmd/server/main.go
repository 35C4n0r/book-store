package main

import (
	"balkan/internal/auth"
	"balkan/internal/database"
	"balkan/internal/handler"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Sanity Test
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"sanity_check": "passed"})
	})

	r.POST("/register", handler.RegisterHandler)
	r.POST("/login", handler.LoginHandler)

	authorized := r.Group("/")
	authorized.Use(auth.AuthenticationMiddleware())

	authorized.GET("/middleware", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"middleware_check": "passed"})
	})

	authorized.POST("/deactivate", handler.DeactivateUserAccountHandler)
	authorized.DELETE("/delete", handler.DeleteUserAccountHandler)

	authorized.POST("/sudo", handler.RegisterAdminHandler)

	authorized.GET("/books", handler.SearchBookHandler)

	authorized.POST("/cart", handler.AddToCartHandler)
	authorized.PUT("/cart", handler.UpdateCartItemHandler)
	authorized.DELETE("/cart/:isbn", handler.DeleteCartHandler)
	authorized.GET("/cart", handler.GetCartHandler)

	authorized.POST("/purchase", handler.PurchaseBookHandler)

	authorized.POST("/review", handler.AddReviewHandler)
	authorized.GET("/review", handler.GetReviewHandler)

	authorized.GET("/download/:isbn", handler.DownloadBookHandler)

	admin := authorized.Group("/admin")
	admin.Use(auth.AdminMiddleware())
	admin.POST("/add", handler.AddBookHandler)
	admin.DELETE("/delete/:isbn", handler.DeleteBookHandler)
	admin.PUT("/update", handler.UpdateBookHandler)

	return r
}

func main() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}

	logrus.SetOutput(&lumberjack.Logger{
		Filename:   "D:\\Jay\\BalkanId\\vit-2025-summer-engineering-internship-task-35C4n0r\\log\\my_app.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})

	// Initializes the Database
	database.InitDb()

	r := setupRouter()
	err := r.Run(":" + os.Getenv("SERVER_PORT"))
	if err != nil {
		logrus.Fatalf("Failed to start the server, error: %v", err)
		return
	}
}
