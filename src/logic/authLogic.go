package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"maily/go-backend/src/models"
	"maily/go-backend/src/utils/token"
	"strings"
)

func GetUserByID(c *gin.Context, uid uint) (models.User, error) {
	db := c.MustGet("DB").(*gorm.DB)

	var user models.User
	if err := db.First(&user, uid).Error; err != nil {
		return user, fmt.Errorf("User not found!")
	}
	user.Password = ""

	return user, nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(c *gin.Context, inputUser models.User) (string, error) {
	db := c.MustGet("DB").(*gorm.DB)

	storedUser := models.User{}
	db.Model(models.User{}).Where("email = ?", inputUser.Email).Take(&storedUser)

	err := VerifyPassword(inputUser.Password, storedUser.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	generateToken, tokenError := token.GenerateToken(storedUser.ID)
	if tokenError != nil {
		return "", tokenError
	}

	return generateToken, nil
}

func SaveUser(c *gin.Context, user models.User) (models.User, error) {
	db := c.MustGet("DB").(*gorm.DB)
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
