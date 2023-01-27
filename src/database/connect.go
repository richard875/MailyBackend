package database

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"maily/go-backend/src/models"
	"os"
)

func Connect() gin.HandlerFunc {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=true", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = DB.AutoMigrate(
		&models.Record{},
		&models.User{},
		&models.Tracker{},
	)

	if err != nil {
		return nil
	}

	// Continue and Return
	return func(c *gin.Context) {
		c.Set("DB", DB)
		c.Next()
	}
}
