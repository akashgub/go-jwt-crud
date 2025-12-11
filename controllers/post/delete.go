package post

import (
	"go-jwt-crud/database"
	"go-jwt-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	//check if posst exists
	result := database.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	// Delete the post
	deleteResult := database.DB.Delete(&post)
	if deleteResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete posst",
		})
		return
	}
	//database.DB.Delete(&models.Post{}, id)

	c.JSON(http.StatusOK, gin.H{"message": "Post Deleted successfully"})
}
