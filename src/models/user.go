package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email             string `gorm:"size:255;not null;unique" json:"email"`
	Password          string `gorm:"size:255;not null;" json:"password"`
	EmailVerification bool   `gorm:"default:true" json:"emailVerification"`
}
