package models

import (
	"gorm.io/gorm"
)

type Tracker struct {
	gorm.Model
	ID                    string `gorm:"size:255;not null;unique;primary_key" json:"id"`        // public tracking number
	PrivateTrackingNumber string `gorm:"size:255;not null;unique" json:"privateTrackingNumber"` // private tracking number
	UserID                string `gorm:"type:char(36); foreignKey:UserRefer" json:"userId"`
}
