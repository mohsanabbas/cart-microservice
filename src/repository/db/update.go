package db

import (
	"github.com/mohsanabbas/cart-microservice/src/domain/cart"
	"github.com/mohsanabbas/cart-microservice/src/util/rest_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
