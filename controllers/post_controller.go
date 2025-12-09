package controllers

import (
	"go-jwt-crud/database"
	"go-jwt-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var post models.Post
	c.ShouldBindJSON(&post)

	database.DB.Create(&post)
	c.JSON(http.StatusOK, post)
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	database.DB.Find(&posts)
	c.JSON(http.StatusOK, posts)
}

func GetPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")

	database.DB.First(&post, id)
	c.JSON(http.StatusOK, post)
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")

	database.DB.First(&post, id)
	c.ShouldBindJSON(&post)
	database.DB.Save(&post)

	c.JSON(http.StatusOK, post)
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Post{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
