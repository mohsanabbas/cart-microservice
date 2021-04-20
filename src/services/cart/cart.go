package cart

import (
	"github.com/mohsanabbas/cart-microservice/src/domain/cart"
	"github.com/mohsanabbas/cart-microservice/src/repository/db"
	"github.com/mohsanabbas/ticketing_utils-go/rest_errors"
)

type Service interface {
	Create(cart.Items) (*cart.Cart, rest_errors.RestErr)
}

type service struct {
	dbRepo db.DbRepository
}

func NewService(dbRepo db.DbRepository) Service {
	return &service{
		dbRepo: dbRepo,
	}
}

func (s *service) Create(request cart.Items) (*cart.Cart, rest_errors.RestErr) {
	cart, err := s.dbRepo.Create(request)
	if err != nil {
		return nil, err
	}

	return cart, nil
}
