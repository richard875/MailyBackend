package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"maily/go-backend/src/models"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() gin.HandlerFunc {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = DB.AutoMigrate(&models.Record{}, &models.User{})
	if err != nil {
		return nil
	}

	return func(c *gin.Context) {
		c.Set("DB", DB)
		c.Next()
	}
}
