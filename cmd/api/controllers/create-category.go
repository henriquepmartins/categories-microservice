package controllers

import (
	"github.com/gin-gonic/gin"
	useCases "github.com/henriquepermartins/categories-ms/internal/use-cases"
)

type createCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(context *gin.Context) {
	useCases.NewCreateCategoryUseCase().Execute(context.Param("name"))

}
