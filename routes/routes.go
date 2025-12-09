package routes

import (
	"go-jwt-crud/controllers"
	"go-jwt-crud/controllers/post"
	"go-jwt-crud/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	api := r.Group("/api")

	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	posts := api.Group("/posts")
	posts.Use(middlewares.JWTAuth())
	{
		posts.POST("", post.CreatePost)
		posts.GET("", post.GetPosts)
		posts.GET("/:id", post.GetPost)
		posts.PUT("/:id", post.UpdatePost)
		posts.DELETE("/:id", post.DeletePost)
	}
}
