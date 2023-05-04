package utils

import (
	"context"
	"errors"
	"jenisha/e-market/models"
	"jenisha/e-market/services"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderServiceImpl struct {
	OrderCollection *mongo.Collection
	ctx             context.Context
}

func NewOrderService(OrderCollection *mongo.Collection, ctx context.Context) services.OrderService {
	return &OrderServiceImpl{
		OrderCollection: OrderCollection,
		ctx:             ctx}
}

func (os *OrderServiceImpl) CreateOrder(order *models.CreateOrder) (*models.DBOrder, error) {
	order.CreatedAt = time.Now()
	order.UpdatedAt = order.CreatedAt
	res, err := os.OrderCollection.InsertOne(os.ctx, order)
	if err != nil {
		return nil, err
	}
	var newOrder *models.DBOrder
	query := bson.M{"_id": res.InsertedID}
	if err = os.OrderCollection.FindOne(os.ctx, query).Decode(&newOrder); err != nil {
		return nil, err
	}

	return newOrder, nil
}

func (os *OrderServiceImpl) FindOrderByID(id string) (*models.DBOrder, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var order *models.DBOrder
	if err := os.OrderCollection.FindOne(os.ctx, query).Decode(&order); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("No document with that Id exists")
		}

		return nil, err
	}
	return order, nil

}

func (os *OrderServiceImpl) FindOrders() ([]*models.DBOrder, error) {
	query := bson.M{}
	var orders []*models.DBOrder

	cursor, err := os.OrderCollection.Find(os.ctx, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(os.ctx)
	for cursor.Next(os.ctx) {
		order := &models.DBOrder{}
		err := cursor.Decode(order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (os *OrderServiceImpl) DeleteOrder(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := os.OrderCollection.DeleteOne(os.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no document with that Id exists")
	}

	return nil
}

func (os *OrderServiceImpl) UpdateOrder(id string, order *models.UpdateOrder) (*models.DBOrder, error) {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	order.UpdatedAt = time.Now()

	if order.OrderStatus == "Dispatched" {
		order.DispatchDate = order.UpdatedAt
	}

	// Reference: https://stackoverflow.com/a/59424834
	objByte, err := bson.Marshal(order) // encoding to bytes
	if err != nil {
		return nil, err
	}

	var update bson.M
	err = bson.Unmarshal(objByte, &update) // decoding to a BSON object
	if err != nil {
		return nil, err
	}

	_, err = os.OrderCollection.UpdateOne(os.ctx, query, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		return nil, err
	}
	var updatedOrder *models.DBOrder
	if err = os.OrderCollection.FindOne(os.ctx, query).Decode(&updatedOrder); err != nil {
		return nil, err
	}
	return updatedOrder, nil
}
