package post

import (
	"go-jwt-crud/database"
	"go-jwt-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//users
// orders/{id}/cannceld

func CreatePost(c *gin.Context) {
	var post models.Post

	//Validata request body
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ //400 → Bad Request
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}
	// save post to DB
	if err := database.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create post",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{ //201 → Created
		"message": "post created successfully",
		"data":    post,
	})
}
