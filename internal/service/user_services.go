package service

import (
	"balkan/internal/database"
	"balkan/internal/models"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func GetUserByEmail(Email string) (*models.UserCRUD, error) {
	var dbUser models.UserCRUD
	result := database.DB.Where("email = ?", Email).First(&dbUser)
	if result.Error != nil {
		logrus.Errorf("Error while getting user by email, %v", result.Error)
		return nil, result.Error
	}
	return &dbUser, nil
}

func GetUserByID(UserId string) (*models.UserCRUD, error) {
	var dbUser models.UserCRUD
	result := database.DB.Where("id = ?", UserId).First(&dbUser)
	if result.Error != nil {
		logrus.Errorf("Error while getting user by email, %v", result.Error)
		return nil, result.Error
	}
	return &dbUser, nil
}

func CreateUser(user models.UserCRUD) error {
	var result = database.DB.Create(&user)
	if result.Error != nil {
		logrus.Errorf("Error while creating a new user, %v", result.Error)
		return result.Error
	}
	return nil
}

func VerifyUser(Email string, Password string) (*models.UserCRUD, error) {
	var dbUser, err = GetUserByEmail(Email)
	if err != nil {
		logrus.Errorf("Error while Verifying user, %v", err)
		return nil, errors.New("INVALID USERNAME/PASSWORD")
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.PasswordHash), []byte(Password))
	if err != nil {
		logrus.Errorf("Error while comparing hash and password, %v", err)
		return nil, errors.New("INVALID USERNAME/PASSWORD")
	}
	return dbUser, nil
}

func DeactivateUser(UserId string) error {
	var dbUser models.UserCRUD
	fmt.Println("UserID:", UserId)
	var result = database.DB.Model(&dbUser).Where("id = ?", UserId).Update("is_active", false)
	if result.Error != nil {
		logrus.Errorf("Error while Deactivating user, %v", result.Error)
		return result.Error
	}
	return nil
}

func DeleteUser(UserId string) error {
	var dbUser models.UserCRUD
	var result = database.DB.Debug().Where("id = ?", UserId).Delete(&dbUser)
	if result.Error != nil {
		logrus.Errorf("Error while Deleting user, %v", result.Error)
		return result.Error
	}
	return nil
}

func RegisterAdmin(UserId string) error {
	var dbUser models.UserCRUD
	result := database.DB.Model(&dbUser).Where("id = ?", UserId).Update("is_admin", true)
	if result.Error != nil {
		logrus.Errorf("Error while registering admin, %v", result.Error)
		return result.Error
	}
	return nil
}
