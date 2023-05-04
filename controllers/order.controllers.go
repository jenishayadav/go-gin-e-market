package controllers

import (
	"fmt"
	"jenisha/e-market/models"
	"jenisha/e-market/services"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService   services.OrderService
	productService services.ProductService
}

// Initializing Controller
func NewOrderController(orderService services.OrderService, productService services.ProductService) OrderController {
	return OrderController{
		orderService:   orderService,
		productService: productService,
	}
}

func (oc OrderController) CreateOrder(ctx *gin.Context) {
	// Note: Can be moved to some constants file
	const (
		maxOrderQty        int64   = 10
		premiumItemsThresh int32   = 3
		premiumDiscount    float32 = 10.0 / 100.0
	)

	var ord *models.CreateOrder
	if err := ctx.ShouldBindJSON(&ord); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var totalValue float32 = 0
	var premiumCount int32 = 0
	products := []*models.UpdateProduct{}
	for _, item := range ord.OrderItems {
		// Check if the quantity is not exceeding maxOrderQty (10)
		if item.Quantity > maxOrderQty {
			ctx.JSON(
				http.StatusBadRequest, gin.H{
					"status":  "fail",
					"message": fmt.Sprintf("Order items Qty should be less than equals to %d", maxOrderQty),
				})
			return
		}
		product_id := item.ProductID
		prod, err := oc.productService.FindProductByID(product_id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": fmt.Sprintf("No product returned with given id %s", product_id),
			})
			return
		}
		// Check if ordered Qty is not exceeding availableQty
		if prod.AvailableQty < item.Quantity {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": fmt.Sprintf("The order Qty is greater than availableQty of product with id %s", product_id),
			})
			return
		}
		totalValue += float32(item.Quantity) * prod.Price
		if prod.Category == "Premium" {
			premiumCount += 1
		}
		// Modify the products by subtracting ordered Qty from available Qty
		// NOTE: Since we need to call productService.UpdateProduct function
		// 	 and that function takes input as object of models.UpdateProduct
		//	 we will need to create object of (ptr) models.UpdateProduct from "prod" which is of models.DBProduct
		var updateProd *models.UpdateProduct = &models.UpdateProduct{
			Id:           prod.Id,
			AvailableQty: prod.AvailableQty - item.Quantity,
		}
		products = append(products, updateProd)
	}

	// Give premium discount on the basis of number of different premium items
	if premiumCount >= premiumItemsThresh {
		totalValue -= totalValue * premiumDiscount
	}

	ord.OrderValue = totalValue
	ord.CreatedAt = time.Now()
	ord.UpdatedAt = ord.CreatedAt
	ord.OrderStatus = "Placed"

	newOrd, err := oc.orderService.CreateOrder(ord)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// TODO
	// Update the products in DB with updated available Qty
	for _, product := range products {
		id := product.Id.Hex()
		_, err = oc.productService.UpdateProduct(id, product)
		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newOrd})
}

func (oc OrderController) FindOrderByID(ctx *gin.Context) {
	id := ctx.Param("id")
	ord, err := oc.orderService.FindOrderByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": ord})
}

func (oc OrderController) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	err := oc.orderService.DeleteOrder(id)
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

func (oc OrderController) FindOrders(ctx *gin.Context) {
	orders, err := oc.orderService.FindOrders()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "Success", "data": orders})

}

func (oc OrderController) UpdateOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	var order *models.UpdateOrder
	// Doubt: Shouldn't we limit only orderStatus to be modified
	// And also orderStatus cycles
	// - only allow status change
	//		(1) from "Placed" to "Dispatched",
	//		(2) from "Dispatched" to "Completed"
	// 		(3) or "Cancelled" anytime.
	//     Meaning, you cannot go "Placed" to "Completed", etc.

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	updatedOrder, err := oc.orderService.UpdateOrder(id, order)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedOrder})
}
