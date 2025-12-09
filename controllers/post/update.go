package post

import (
	"go-jwt-crud/database"
	"go-jwt-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdatePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")

	database.DB.First(&post, id)
	c.ShouldBindJSON(&post)
	database.DB.Save(&post)

	c.JSON(http.StatusOK, post)
}
