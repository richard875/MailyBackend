package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"html"
	"maily/go-backend/src/database"
	"maily/go-backend/src/models"
	"maily/go-backend/src/utils/token"
	"net/mail"
	"strings"
)

func GetUserByID(c *gin.Context, uid string) (models.User, error) {
	db := database.DB

	var user models.User
	if err := db.First(&user, "id = ?", uid).Error; err != nil {
		return user, fmt.Errorf("user not found")
	}
	user.Password = "hidden"

	return user, nil
}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(inputUser models.User) (string, error) {
	db := database.DB

	storedUser := models.User{}
	result := db.First(&storedUser, "email = ?", inputUser.Email)

	if result.Error != nil {
		return "", result.Error
	}

	err := verifyPassword(inputUser.Password, storedUser.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	generateToken, tokenError := token.GenerateToken(storedUser.ID)
	if tokenError != nil {
		return "", tokenError
	}

	return generateToken, nil
}

func SaveUser(user models.User) (models.User, error) {
	db := database.DB
	result := db.Create(&user)

	return user, result.Error
}

func HashPassword(password string) string {
	// Turn password into hash
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func TrimTextAndVerifyEmail(c *gin.Context, username string) (string, error) {
	// Remove spaces in username
	email := strings.TrimSpace(username)
	_, err := mail.ParseAddress(email)

	return html.EscapeString(email), err
}
