package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	useCase "github.com/henriquepermartins/categories-ms/internal/use-cases"
)

type createCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(context *gin.Context) {
	var body createCategoryInput

	if err := context.ShouldBindJSON(&body); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"sucess": false, "error": err.Error()})
		return
	}

	useCase := useCase.NewCreateCategoryUseCase()

	err := useCase.Execute(body.Name)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"sucess": false, "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"sucess": true})
}
