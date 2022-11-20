package models

type Register struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
