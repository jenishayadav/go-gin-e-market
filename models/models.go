package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ProductID string `json:"productId" bson:"productId"`
	Quantity  int64  `json:"quantity" bson:"quantity" binding:"required"`
}

type CreateOrder struct {
	OrderItems   []OrderItem `json:"orderItems,omitempty" bson:"orderItems,omitempty" `
	OrderValue   float32     `json:"orderValue,omitempty" bson:"orderValue,omitempty"`
	DispatchDate time.Time   `json:"dispatchDate,omitempty" bson:"dispatchDate,omitempty"`
	OrderStatus  string      `json:"orderStatus,omitempty" bson:"orderStatus,omitempty"`
	CreatedAt    time.Time   `json:"createdAt,omitempty" bson:"createdAt,omitempty" `
	UpdatedAt    time.Time   `json:"updatedAt,omitempty" bson:"updatedAt,omitempty" `
}

type DBOrder struct {
	Id           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	OrderItems   []OrderItem        `json:"orderItems,omitempty" bson:"orderItems,omitempty" `
	OrderValue   float32            `json:"orderValue" bson:"orderValue" binding:"required"`
	DispatchDate time.Time          `json:"dispatchDate" bson:"dispatchDate" binding:"omitempty"`
	OrderStatus  string             `json:"orderStatus" bson:"orderStatus" binding:"required"`
	CreatedAt    time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty" `
	UpdatedAt    time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty" `
}

type UpdateOrder struct {
	Id           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	OrderItems   []OrderItem        `json:"orderItems,omitempty" bson:"orderItems,omitempty"`
	OrderValue   float32            `json:"orderValue,omitempty" bson:"orderValue,omitempty"`
	DispatchDate time.Time          `json:"dispatchDate,omitempty" bson:"dispatchDate,omitempty"`
	OrderStatus  string             `json:"orderStatus,omitempty" bson:"orderStatus,omitempty"`
	CreatedAt    time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt    time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type CreateProduct struct {
	Name         string    `json:"name" bson:"name" binding:"required"`
	AvailableQty int64     `json:"availableQty" bson:"availableQty" binding:"required"`
	Price        float32   `json:"price" bson:"price" binding:"required"`
	Category     string    `json:"category" bson:"category" binding:"required"`
	CreatedAt    time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt    time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type DBProduct struct {
	Id           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name" binding:"required"`
	AvailableQty int64              `json:"availableQty" bson:"availableQty" binding:"required"`
	Price        float32            `json:"price" bson:"price" binding:"required"`
	Category     string             `json:"category" bson:"category" binding:"required"`
	CreatedAt    time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty" `
	UpdatedAt    time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty" `
}

type UpdateProduct struct {
	Id           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name,omitempty"`
	AvailableQty int64              `json:"availableQty,omitempty" bson:"availableQty,omitempty"`
	Price        float32            `json:"price,omitempty" bson:"price,omitempty"`
	Category     string             `json:"category,omitempty" bson:"category,omitempty"`
	CreatedAt    time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt    time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
