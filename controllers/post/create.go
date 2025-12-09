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
	c.ShouldBindJSON(&post)

	database.DB.Create(&post)
	c.JSON(http.StatusOK, post)
}
