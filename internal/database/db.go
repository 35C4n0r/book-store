package database

import (
	"balkan/internal/models"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	// your models package here, for example:
	// "your_project/internal/models"
)

var DB *gorm.DB

func InitDb() {
	var err error
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	timeZone := os.Getenv("TIME_ZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, timeZone)
	//dsn := "host=core_postgresql user=root password=password dbname=bookstore port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.Fatalf("Error connecting to database: %v", err)
	} else {
		logrus.Infof("Database Connected")
	}

	// AutoMigrate your models here, for example:
	err = DB.AutoMigrate(&models.UserCRUD{})
	err = DB.AutoMigrate(&models.InventoryCRUD{})

	err = DB.AutoMigrate(&models.CartCRUD{})
	//DB.Exec("ALTER TABLE cart ADD CONSTRAINT fk_inventory FOREIGN KEY (isbn) REFERENCES inventory(isbn) ON DELETE SET NULL ON UPDATE CASCADE;")

	err = DB.AutoMigrate(&models.ReviewCRUD{})
	err = DB.AutoMigrate(&models.PurchaseCRUD{})
	//DB.Exec("ALTER TABLE reviews ADD CONSTRAINT fk_reviews FOREIGN KEY (isbn) REFERENCES reviews(isbn) ON DELETE SET NULL ON UPDATE CASCADE;")

	if err != nil {
		logrus.Fatalf("Error migrating models: %v", err)
		return
	}
	// DB.AutoMigrate(&models.Book{})
}

// CloseDb closes the database connection
func CloseDb() {
	db, err := DB.DB()
	if err != nil {
		logrus.Fatalf("Could not get database instance: %v", err)
	}
	err = db.Close()
	if err != nil {
		return
	}
}
