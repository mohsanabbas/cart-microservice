package db

import (
	"github.com/mohsanabbas/cart-microservice/src/domain/cart"
	"github.com/mohsanabbas/cart-microservice/src/util/rest_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *dbRepository) Create(doc cart.Cart) (*cart.Cart, rest_errors.RestErr) {
	res, err := r.col.InsertOne(r.ctx, doc)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("Error while saving items in database", err)
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return r.GetById(id)
}
