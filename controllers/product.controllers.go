package controllers

import (
	"jenisha/e-market/models"
	"jenisha/e-market/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService services.ProductService
}

// Initializing Controller

func NewProductController(productService services.ProductService) ProductController {
	return ProductController{productService: productService}
}

func (pc ProductController) CreateProduct(ctx *gin.Context) {
	var prod *models.CreateProduct

	if err := ctx.ShouldBindJSON(&prod); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newProd, err := pc.productService.CreateProduct(prod)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newProd})
}

func (pc ProductController) FindProductByID(ctx *gin.Context) {
	id := ctx.Param("id")
	prod, err := pc.productService.FindProductByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": prod})
}

func (pc ProductController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	err := pc.productService.DeleteProduct(id)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)

}

func (pc ProductController) FindProducts(ctx *gin.Context) {
	products, err := pc.productService.FindProducts()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "Success", "data": products})

}

func (pc ProductController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	var product *models.UpdateProduct
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	updatedProd, err := pc.productService.UpdateProduct(id, product)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedProd})
}
