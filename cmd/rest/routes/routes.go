package rest

import (
	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	v1:= router.Group("/v1") 
	return v1
}