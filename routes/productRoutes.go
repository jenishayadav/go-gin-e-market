package routes

import (
	"jenisha/e-market/controllers"

	"github.com/gin-gonic/gin"
)

type ProductRoutesController struct {
	ProductController controllers.ProductController
}

func NewProductRoute(ProductController controllers.ProductController) ProductRoutesController {
	return ProductRoutesController{
		ProductController: ProductController,
	}
}

func (prod *ProductRoutesController) ProductRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/products")
	router.GET("/", prod.ProductController.FindProducts)
	router.GET("/:id", prod.ProductController.FindProductByID)
	router.POST("/", prod.ProductController.CreateProduct)
	router.DELETE("/:id", prod.ProductController.DeleteProduct)
	router.PATCH("/:id", prod.ProductController.UpdateProduct)
}
