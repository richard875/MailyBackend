package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"maily/go-backend/src/dtos"
	"maily/go-backend/src/utils/token"
	"net/http"

	"maily/go-backend/src/logic"
	"maily/go-backend/src/models"
	"maily/go-backend/src/utils"
)

func CurrentUser(c *gin.Context) {
	userId, tokenError := token.ExtractUserID(c)
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
	var input dtos.Login
	_ = c.ShouldBindJSON(&input)

	user := models.User{}
	user.Email = input.Email
	user.Password = input.Password

	loginCheck, err := logic.LoginCheck(c, user)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": loginCheck})
}

func Register(c *gin.Context) {
	var input dtos.Register
	_ = c.ShouldBindJSON(&input)

	// Create User
	parsedEmail, emailError := logic.TrimTextAndVerifyEmail(c, input.Email)
	if emailError != nil {
		utils.HandleError(c, emailError)
		return
	}

	user := models.User{}
	user.ID = uuid.New()
	user.Email = parsedEmail
	user.Password = logic.HashPassword(input.Password)

	_, saveError := logic.SaveUser(c, user)
	if saveError != nil {
		utils.HandleError(c, saveError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}
