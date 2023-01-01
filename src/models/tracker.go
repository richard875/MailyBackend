package models

import (
	"gorm.io/gorm"
)

type Tracker struct {
	gorm.Model
	ID                string `gorm:"size:255;not null;unique;primary_key" json:"id"` // public tracking number
	UserID            string `gorm:"type:char(36);not null;foreignKey:UserRefer" json:"userId"`
	TimesOpened       int    `gorm:"type:int;not null;default:0" json:"timesOpened"`
	ComposeAction     int    `gorm:"type:int;not null;default:-1" json:"composeAction"` // newMessage: 1, reply: 2, replyAll: 3, forward: 4
	Subject           string `gorm:"type:varchar(255);not null" json:"subject"`
	FromAddress       string `gorm:"type:varchar(255);not null" json:"fromAddress"`
	ToAddresses       string `gorm:"type:varchar(255);not null" json:"toAddresses"`
	CcAddresses       string `gorm:"type:varchar(255);not null" json:"ccAddresses"`
	BccAddresses      string `gorm:"type:varchar(255);not null" json:"bccAddresses"`
	ReplyToAddresses  string `gorm:"type:varchar(255);not null" json:"replyToAddresses"`
	InternalMessageID string `gorm:"type:varchar(255);not null" json:"internalMessageID"` // No computational use, just for reference
}
