package db

import (
	"context"

	"github.com/mohsanabbas/cart-microservice/src/domain/cart"
	"github.com/mohsanabbas/cart-microservice/src/util/rest_errors"
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
