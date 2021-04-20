package db

import (
	"encoding/json"
	"errors"

	"github.com/mohsanabbas/cart-microservice/src/domain/cart"
	"github.com/mohsanabbas/ticketing_utils-go/rest_errors"
)

type DbRepository interface {
	Create(cart.Items) (*cart.Cart, rest_errors.RestErr)
}

func NewCartRepository() DbRepository {
	return &dbRepository{}
}

type dbRepository struct {
}

func (r *dbRepository) Create(at cart.Items) (*cart.Cart, rest_errors.RestErr){
	reqBody := cart.Items{
		Items: at.Items,
	}
	payload, err := json.Marshal(reqBody)

	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to marshal json request", errors.New("request error cart-microservice"))
	}

	var result cart.Cart

	if err := json.Unmarshal(payload, &result); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshal cart-microservice response", errors.New("json parsing error"))
	}
	// if err := mongo.Create(); err != nil {
	// 	return rest_errors.NewInternalServerError("error when trying to save access token in database", err)
	// }
	return &result, nil
}
