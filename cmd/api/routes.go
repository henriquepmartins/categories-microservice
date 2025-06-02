package main

import (
	"github.com/gin-gonic/gin"
	"github.com/henriquepermartins/categories-ms/cmd/api/controllers"
	"github.com/henriquepermartins/categories-ms/internal/repositories"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")

	categoryRepository := repositories.NewCategoryRepository()
	
	categoryRoutes.POST("/", func(ctx *gin.Context) {
		controllers.CreateCategory(ctx, categoryRepository)
	})

}
