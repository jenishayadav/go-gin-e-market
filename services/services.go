package services

import "jenisha/e-market/models"

type OrderService interface {
	CreateOrder(*models.CreateOrder) (*models.DBOrder, error)
	FindOrderByID(string) (*models.DBOrder, error)
	FindOrders() ([]*models.DBOrder, error)
	UpdateOrder(string, *models.UpdateOrder) (*models.DBOrder, error)
	DeleteOrder(string) error
}
type ProductService interface {
	CreateProduct(*models.CreateProduct) (*models.DBProduct, error)
	FindProductByID(string) (*models.DBProduct, error)
	FindProducts() ([]*models.DBProduct, error)
	UpdateProduct(string, *models.UpdateProduct) (*models.DBProduct, error)
	// ReplaceProduct(string, *models.UpdateProduct) (*models.DBProduct, error)
	DeleteProduct(string) error
}
