package controllers

import (
	"github.com/gin-gonic/gin"
	"maily/go-backend/src/utils/token"
	"net/http"

	"maily/go-backend/src/logic"
	"maily/go-backend/src/models"
	"maily/go-backend/src/utils"
)

func CurrentUser(c *gin.Context) {
	userId, tokenError := token.ExtractTokenID(c)
	if tokenError != nil {
		utils.HandleError(c, tokenError)
		return
	}

	user, userError := logic.GetUserByID(c, userId)
	if userError != nil {
		utils.HandleError(c, userError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "loginCheck": user})
}

func Login(c *gin.Context) {
	var input models.Login
	_ = c.ShouldBindJSON(&input)

	user := models.User{}
	user.Email = input.Email
	user.Password = input.Password

	loginCheck, err := logic.LoginCheck(c, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": loginCheck})
}

func Register(c *gin.Context) {
	var input models.Register
	_ = c.ShouldBindJSON(&input)

	// Create User
	user := models.User{}
	user.Email = logic.TrimUsername(input.Email)
	user.Password = logic.HashPassword(input.Password)

	_, err := logic.SaveUser(c, user)
	if err != nil {
		utils.HandleError(c, err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}
