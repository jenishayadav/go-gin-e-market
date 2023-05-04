package routes

import (
	"jenisha/e-market/controllers"

	"github.com/gin-gonic/gin"
)

type OrderRoutesController struct {
	OrderController controllers.OrderController
}

func NewOrderRoute(OrderController controllers.OrderController) OrderRoutesController {
	return OrderRoutesController{
		OrderController: OrderController,
	}
}

func (ord *OrderRoutesController) OrderRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/orders")
	router.GET("/", ord.OrderController.FindOrders)
	router.GET("/:id", ord.OrderController.FindOrderByID)
	router.POST("/", ord.OrderController.CreateOrder)
	router.DELETE("/:id", ord.OrderController.DeleteOrder)
	router.PATCH("/:id", ord.OrderController.UpdateOrder)

}
