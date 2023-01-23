package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	ID                       uuid.UUID `gorm:"type:char(36);not null;unique;primary_key" json:"id"`
	PublicTrackingNumber     string    `gorm:"type:char(36);not null;foreignKey:TrackerRefer" json:"publicTrackingNumber"`
	IpAddress                string    `gorm:"size:255" json:"ipAddress"`
	IpCity                   string    `gorm:"size:255" json:"ipCity"`
	IpCountry                string    `gorm:"size:255" json:"ipCountry"`
	EmojiFlag                string    `gorm:"size:255" json:"emojiFlag"`
	IsEU                     bool      `gorm:"default:false" json:"isEu"`
	IsTor                    bool      `gorm:"default:false" json:"isTor"`
	IsProxy                  bool      `gorm:"default:false" json:"isProxy"`
	IsAnonymous              bool      `gorm:"default:false" json:"isAnonymous"`
	IsKnownAttacker          bool      `gorm:"default:false" json:"isKnownAttacker"`
	IsKnownAbuser            bool      `gorm:"default:false" json:"isKnownAbuser"`
	IsThreat                 bool      `gorm:"default:false" json:"isThreat"`
	IsBogon                  bool      `gorm:"default:false" json:"isBogon"`
	Latitude                 float64   `gorm:"type:decimal(10,8)" json:"latitude"`
	Longitude                float64   `gorm:"type:decimal(11,8)" json:"longitude"`
	ConfidentWithEmailClient bool      `gorm:"default:false" json:"confidentWithEmailClient"`
	Headers                  string    `gorm:"type:varchar(5000)"  json:"header"`
}
