package main

import (
	"github.com/gin-gonic/gin"
	"github.com/henriquepermartins/categories-ms/cmd/api/controllers"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")

	categoryRoutes.POST("/", controllers.CreateCategory)

}
