package post

import (
	"go-jwt-crud/database"
	"go-jwt-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
