package main

import (
	"go-jwt-crud/database"
	"go-jwt-crud/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.ConnectDB()

	r := gin.Default()
	routes.Setup(r)

	r.Run(":" + os.Getenv("PORT"))
}
