package cart

import (
	"strings"

	"github.com/mohsanabbas/cart-microservice/src/domain/cart"
	"github.com/mohsanabbas/cart-microservice/src/repository/db"
	"github.com/mohsanabbas/ticketing_utils-go/rest_errors"
)

type Service interface {
	Create(cart.Cart) (*cart.Cart, rest_errors.RestErr)
	GetById(string) (*cart.Cart, rest_errors.RestErr)
}

type service struct {
	dbRepo db.DbRepository
}

func NewService(dbRepo db.DbRepository) Service {
	return &service{
		dbRepo: dbRepo,
	}
}

func (s *service) Create(request cart.Cart) (*cart.Cart, rest_errors.RestErr) {

	request.SetCartExpiration()
	cart, err := s.dbRepo.Create(request)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (s *service) GetById(id string) (*cart.Cart, rest_errors.RestErr) {
	id = strings.TrimSpace(id)
	if len(id) == 0 {
		return nil, rest_errors.NewBadRequestError("invalid cart id")
	}
	cart, err := s.dbRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return cart, nil
}
