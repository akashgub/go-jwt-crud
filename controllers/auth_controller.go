package controllers

import (
	"go-jwt-crud/database"
	"go-jwt-crud/models"
	"go-jwt-crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	c.ShouldBindJSON(&user)

	user.Password = utils.HashPassword(user.Password)
	database.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}

func Login(c *gin.Context) {
	var input models.User
	var user models.User

	c.ShouldBindJSON(&input)

	database.DB.Where("email = ?", input.Email).First(&user)
	if user.ID == 0 || !utils.CheckPassword(user.Password, input.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, _ := utils.GenerateToken(user.ID)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
