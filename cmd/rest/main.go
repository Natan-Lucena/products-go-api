package main

import (
	rest "github.com/Natan-Lucena/products-api/cmd/rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	rest.AppRoutes(app)
	app.Run("localhost:3001")
}