package models

import (
	"gorm.io/gorm"
	"time"
)

type Record struct {
	gorm.Model
	UserID          string `gorm:"primaryKey"`
	LogNumber       string
	LogTimes        int `gorm:"default:0"`
	IpAddress       string
	IpCity          string
	IsTor           bool
	IsProxy         bool
	IsAnonymous     bool
	IsKnownAttacker bool
	IsKnownAbuser   bool
	IsThreat        bool
	IsBogon         bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
