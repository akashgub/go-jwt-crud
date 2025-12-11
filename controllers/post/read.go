package post

import (
	"go-jwt-crud/database"
	"go-jwt-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPosts returns all posts
func GetPosts(c *gin.Context) {
	var posts []models.Post

	if err := database.DB.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch posts",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": posts})
}

// GetPost returns a single post by ID
func GetPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")

	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}
