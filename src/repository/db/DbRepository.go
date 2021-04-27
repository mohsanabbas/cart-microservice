package db

import (
	"context"
	"errors"

	"github.com/mohsanabbas/cart-microservice/src/domain/cart"
	"github.com/mohsanabbas/cart-microservice/src/util/rest_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// DbRepository repository interface
type DbRepository interface {
	Create(cart.Cart) (*cart.Cart, rest_errors.RestErr)
	GetById(string) (*cart.Cart, rest_errors.RestErr)
	Update(string, cart.Item) (*cart.CartUpdate, rest_errors.RestErr)
	Delete(string, string) (*cart.CartUpdate, rest_errors.RestErr)
	DeleteAll(string) (*cart.CartUpdate, rest_errors.RestErr)
}

type dbRepository struct {
	ctx context.Context
	col *mongo.Collection
}

// NewCartRepository construct repoitory
func NewCartRepository(ctx context.Context, col *mongo.Collection) DbRepository {
	return &dbRepository{
		ctx: ctx,
		col: col,
	}
}

// Create generate new cart
func (r *dbRepository) Create(doc cart.Cart) (*cart.Cart, rest_errors.RestErr) {
	res, err := r.col.InsertOne(r.ctx, doc)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("Error while saving items in database", err)
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return r.GetById(id)
}

// GetById find cart with "_id"
func (r *dbRepository) GetById(id string) (*cart.Cart, rest_errors.RestErr) {
	cart := cart.Cart{}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error converting ID", errors.New("ID conversion failed"))
	}

	filter := bson.M{"_id": _id}
	if err := r.col.FindOne(r.ctx, filter).Decode(&cart); err != nil {
		return nil, rest_errors.NewNotFoundError("Cart not found in database")
	}

	return &cart, nil
}

// Update insert item in existing cart by cartid
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
		return nil, rest_errors.NewInternalServerError("cart not found", err)
	}
	result.ModifiedCount = res.ModifiedCount
	result.Results = *updCart
	return &result, nil
}

// Delete item in cart by "_id"
func (r *dbRepository) Delete(cartId string, itemId string) (*cart.CartUpdate, rest_errors.RestErr) {
	result := cart.CartUpdate{
		ModifiedCount: 0,
	}
	_cartId, err := primitive.ObjectIDFromHex(cartId)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("Error while converting string id to `ObjectID`", err)
	}
	_itemId, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("Error while converting string id to `ObjectID`", err)
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

// DeleteAll clear cart items
func (r *dbRepository) DeleteAll(cartId string) (*cart.CartUpdate, rest_errors.RestErr) {
	result := cart.CartUpdate{
		ModifiedCount: 0,
	}
	_cartId, err := primitive.ObjectIDFromHex(cartId)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("Error while converting string id to `ObjectID`", err)
	}
	clear := make([]interface{}, 0)
	filter := bson.M{"$set": bson.M{"items": clear}}
	res, err := r.col.UpdateOne(r.ctx, bson.M{"_id": _cartId}, filter)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("Error while clearing cart", err)
	}
	updCart, err := r.GetById(cartId)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("cart not found", err)
	}
	result.ModifiedCount = res.ModifiedCount
	result.Results = *updCart
	return &result, nil
}
