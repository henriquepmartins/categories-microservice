package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henriquepermartins/categories-ms/internal/repositories"
	useCase "github.com/henriquepermartins/categories-ms/internal/use-cases"
)

func ListCategories(context *gin.Context, repository repositories.CategoryRepositoryInterface) {
	useCase := useCase.NewListCategoryUseCase(repository)

	categories, err := useCase.Execute()

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"sucess": false, "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"sucess": true, "data": categories})
}
