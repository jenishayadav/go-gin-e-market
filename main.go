package main

import (
	"context"
	"fmt"
	"jenisha/e-market/controllers"
	"jenisha/e-market/routes"
	"jenisha/e-market/services"
	"jenisha/e-market/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	ctx         context.Context
	mongoClient *mongo.Client

	authCollection *mongo.Collection

	// Add the Order Services, Controllers and Routes
	orderService          services.OrderService
	orderController       controllers.OrderController
	orderCollection       *mongo.Collection
	orderRoutesController routes.OrderRoutesController

	// Add Product Services, Controllers and Routes
	productService          services.ProductService
	productController       controllers.ProductController
	productCollection       *mongo.Collection
	productRoutesController routes.ProductRoutesController
)

func init() {
	// Connect to MongoDB
	ctx = context.TODO()

	// Note: Can be moved to some constants file
	const (
		mongoURI              string = "mongodb://localhost:27017"
		database              string = "josh_assignment"
		orderCollectionName   string = "order"
		productCollectionName string = "product"
	)

	mongoConn := options.Client().ApplyURI(mongoURI)
	mongoClient, err := mongo.Connect(ctx, mongoConn)
	if err != nil {
		panic(err)
	}

	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	productCollection = mongoClient.Database(database).Collection(productCollectionName)
	productService = utils.NewProductService(productCollection, ctx)
	productController = controllers.NewProductController(productService)
	productRoutesController = routes.NewProductRoute(productController)

	orderCollection = mongoClient.Database(database).Collection(orderCollectionName)
	orderService = utils.NewOrderService(orderCollection, ctx)
	orderController = controllers.NewOrderController(orderService, productService)
	orderRoutesController = routes.NewOrderRoute(orderController)

	server = gin.Default()
}

func main() {
	const serverPort string = "8080"

	router := server.Group("/api")

	// Defining healthcheck
	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	orderRoutesController.OrderRoutes(router)
	productRoutesController.ProductRoutes(router)

	log.Fatalln(server.Run(":" + serverPort))
	defer mongoClient.Disconnect(ctx)
}
