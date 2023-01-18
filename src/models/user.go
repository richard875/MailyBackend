package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:char(36);not null;unique;primary_key" json:"id"`
	FirstName     string    `gorm:"type:varchar(255);not null" json:"firstName"`
	LastName      string    `gorm:"type:varchar(255);not null" json:"lastName"`
	Email         string    `gorm:"size:255;not null;unique" json:"email"`
	Password      string    `gorm:"size:255;not null" json:"password"`
	EmailVerified bool      `gorm:"default:true" json:"emailVerified"`
	TotalClicks   int       `gorm:"default:0" json:"totalClicks"`
	EmailsSent    int       `gorm:"default:0" json:"emailsSent"`
	TelegramToken string    `gorm:"size:255" json:"telegramToken"`
	TelegramID    int       `gorm:"default:null" json:"telegramID"`
}
