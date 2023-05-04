package utils

import (
	"context"
	"errors"
	"jenisha/e-market/models"
	"jenisha/e-market/services"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductServiceImpl struct {
	productCollection *mongo.Collection
	ctx               context.Context
}

func NewProductService(ProductCollection *mongo.Collection, ctx context.Context) services.ProductService {
	return &ProductServiceImpl{ProductCollection, ctx}

}

func (ps *ProductServiceImpl) CreateProduct(product *models.CreateProduct) (*models.DBProduct, error) {
	product.CreatedAt = time.Now()
	product.UpdatedAt = product.CreatedAt

	res, err := ps.productCollection.InsertOne(ps.ctx, product)
	if err != nil {
		log.Println("Mongo: Product collection insert one error.")
		return nil, err
	}
	query := bson.M{"_id": res.InsertedID}

	var newProduct *models.DBProduct
	if err = ps.productCollection.FindOne(ps.ctx, query).Decode(&newProduct); err != nil {
		return nil, err
	}
	return newProduct, nil
}

func (ps *ProductServiceImpl) FindProductByID(id string) (*models.DBProduct, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var product *models.DBProduct
	if err := ps.productCollection.FindOne(ps.ctx, query).Decode(&product); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("No document with that Id exists")
		}

		return nil, err
	}
	return product, nil

}

func (ps *ProductServiceImpl) FindProducts() ([]*models.DBProduct, error) {
	query := bson.M{}
	var products []*models.DBProduct

	cursor, err := ps.productCollection.Find(ps.ctx, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ps.ctx)
	for cursor.Next(ps.ctx) {
		product := &models.DBProduct{}
		err := cursor.Decode(product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)

	}
	return products, nil
}

func (ps *ProductServiceImpl) DeleteProduct(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := ps.productCollection.DeleteOne(ps.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no document with that Id exists")
	}

	return nil
}

// func (ps *ProductServiceImpl) ReplaceProduct(id string, product *models.DBProduct) (*models.DBProduct, error) {
// 	obId, _ := primitive.ObjectIDFromHex(id)
// 	query := bson.M{"_id": obId}

// 	product.UpdatedAt = time.Now()

// 	_, err := ps.productCollection.ReplaceOne(ps.ctx, query, product)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var updatedProd *models.DBProduct
// 	if err = ps.productCollection.FindOne(ps.ctx, query).Decode(&updatedProd); err != nil {
// 		return nil, err
// 	}
// 	return updatedProd, nil
// }

func (ps *ProductServiceImpl) UpdateProduct(id string, product *models.UpdateProduct) (*models.DBProduct, error) {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	product.UpdatedAt = time.Now()

	// Reference: https://stackoverflow.com/a/59424834
	objByte, err := bson.Marshal(product) // encoding to bytes
	if err != nil {
		return nil, err
	}

	var update bson.M
	err = bson.Unmarshal(objByte, &update) // decoding to a BSON object
	if err != nil {
		return nil, err
	}

	_, err = ps.productCollection.UpdateOne(ps.ctx, query, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		return nil, err
	}
	var updatedProd *models.DBProduct
	if err = ps.productCollection.FindOne(ps.ctx, query).Decode(&updatedProd); err != nil {
		return nil, err
	}
	return updatedProd, nil
}
