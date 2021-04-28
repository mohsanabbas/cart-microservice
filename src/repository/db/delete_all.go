package db

import (
	"github.com/mohsanabbas/cart-microservice/src/domain/cart"
	"github.com/mohsanabbas/cart-microservice/src/util/rest_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
