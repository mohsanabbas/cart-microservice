package db

import (
	"errors"

	"github.com/mohsanabbas/cart-microservice/src/domain/cart"
	"github.com/mohsanabbas/cart-microservice/src/util/rest_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
