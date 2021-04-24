package db

import (
	"context"
	"errors"

	"github.com/mohsanabbas/cart-microservice/src/domain/cart"
	"github.com/mohsanabbas/ticketing_utils-go/rest_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	expirationTime = 24
)

type DbRepository interface {
	Create(cart.Cart) (*cart.Cart, rest_errors.RestErr)
	GetById(string) (*cart.Cart, rest_errors.RestErr)
	Update(string, cart.Item) (*cart.CartUpdate, rest_errors.RestErr)
	Delete(string, string) (*cart.CartUpdate, rest_errors.RestErr)
}

type dbRepository struct {
	col *mongo.Collection
	ctx context.Context
}

func NewCartRepository(col *mongo.Collection, ctx context.Context) DbRepository {
	return &dbRepository{
		col: col,
		ctx: ctx,
	}
}

// Create
func (r *dbRepository) Create(doc cart.Cart) (*cart.Cart, rest_errors.RestErr) {
	res, err := r.col.InsertOne(r.ctx, doc)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("Error while saving items in database", err)
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return r.GetById(id)
}

// Get
func (r *dbRepository) GetById(id string) (*cart.Cart, rest_errors.RestErr) {
	cart := cart.Cart{}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error converting ID", errors.New("ID conversion failed"))
	}

	filter := bson.M{"_id": _id}
	err = r.col.FindOne(r.ctx, filter).Decode(&cart)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error finding cart in database", err)
	}

	return &cart, nil
}

// Update
func (r *dbRepository) Update(id string, update cart.Item) (*cart.CartUpdate, rest_errors.RestErr) {
	result := cart.CartUpdate{
		ModifiedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("Error id", err)
	}
	filter := bson.M{"$push": bson.M{"items": update}}
	res, err := r.col.UpdateOne(r.ctx, bson.M{"_id": _id}, filter)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("Error while updatig cart", err)
	}

	updCart, err := r.GetById(id)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("updated cart not found", err)
	}
	result.ModifiedCount = res.ModifiedCount
	result.Results = *updCart
	return &result, nil
}

// Delete
func (r *dbRepository) Delete(cartId string, itemId string) (*cart.CartUpdate, rest_errors.RestErr) {
	result := cart.CartUpdate{
		ModifiedCount: 0,
	}
	_cartId, err := primitive.ObjectIDFromHex(cartId)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("Error ", err)
	}
	_itemId, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("Error ", err)
	}
	filter := bson.M{"$pull": bson.M{"items": bson.M{"_id": _itemId}}}
	res, err := r.col.UpdateOne(r.ctx, bson.M{"_id": _cartId}, filter)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("Error while deleting item", err)
	}
	updCart, err := r.GetById(cartId)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("cart not found", err)
	}
	if res.ModifiedCount == 0 {
		return nil, rest_errors.NewInternalServerError("Item not found", errors.New("Item does not exist"))
	}
	result.ModifiedCount = res.ModifiedCount
	result.Results = *updCart
	return &result, nil
}
