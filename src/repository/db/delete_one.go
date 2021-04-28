package db

import (
	"errors"

	"github.com/mohsanabbas/cart-microservice/src/domain/cart"
	"github.com/mohsanabbas/cart-microservice/src/util/rest_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
