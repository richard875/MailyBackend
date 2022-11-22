package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:char(36);not null;unique;primary_key" json:"id"`
	Email         string    `gorm:"size:255;not null" json:"email"`
	Password      string    `gorm:"size:255;not null" json:"password"`
	EmailVerified bool      `gorm:"default:true" json:"emailVerified"`
}
