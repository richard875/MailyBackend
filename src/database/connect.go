package database

import (
	"fmt"

	"maily/go-backend/src/models"
	"maily/go-backend/src/secrets"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() gin.HandlerFunc {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", secrets.DbUsername, secrets.DbPassword, secrets.DbHost, secrets.DbPort, secrets.DbDatabase)
	db, connectError := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if connectError != nil {
		panic("failed to connect database")
	}

	migrateError := db.AutoMigrate(&models.Record{})
	if migrateError != nil {
		return nil
	}

	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}
