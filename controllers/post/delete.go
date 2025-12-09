package post

import (
	"go-jwt-crud/database"
	"go-jwt-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Post{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
