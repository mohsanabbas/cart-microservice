package cart

import (
	"github.com/mohsanabbas/cart-microservice/src/domain/cart"
	"github.com/mohsanabbas/cart-microservice/src/repository/db"
	"github.com/mohsanabbas/cart-microservice/src/util/rest_errors"
)

// Usecase interface
type Usecase interface {
	Create(cart.Cart, cart.RequestHeaders) (*cart.Cart, rest_errors.RestErr)
	GetById(string) (*cart.Cart, rest_errors.RestErr)
	Update(string, cart.Item) (*cart.CartUpdate, rest_errors.RestErr)
	Delete(string, string) (*cart.CartUpdate, rest_errors.RestErr)
	DeleteAll(string) (*cart.CartUpdate, rest_errors.RestErr)
}

type usecase struct {
	dbRepo db.DbRepository
}

// NewUsecase construct usecase
func NewUsecase(dbRepo db.DbRepository) Usecase {
	return &usecase{
		dbRepo: dbRepo,
	}
}
