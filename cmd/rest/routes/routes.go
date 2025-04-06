package rest

import (
	"github.com/Natan-Lucena/products-api/internal/infraestructure/factories"
	"github.com/Natan-Lucena/products-api/internal/infraestructure/middlewares"
	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	createProductController := factories.GetCreateProductController()

	v1:= router.Group("/v1") 
	v1.Use(middlewares.AuthMiddleware())
	v1.POST("/products", createProductController.Handle)
	return v1
}