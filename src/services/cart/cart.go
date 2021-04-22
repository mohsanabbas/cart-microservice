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
	Update(string, cart.Item) (*cart.CartUpdate, rest_errors.RestErr)
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
	for k := range request.Items {
		request.Items[k].GenerateItemID()
	}
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

func (s *service) Update(id string, request cart.Item) (*cart.CartUpdate, rest_errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	id = strings.TrimSpace(id)
	if len(id) == 0 {
		return nil, rest_errors.NewBadRequestError("invalid cart id")
	}
	request.GenerateItemID()

	cart, err := s.dbRepo.Update(id, request)
	if err != nil {
		return nil, err
	}
	return &cart, nil
}
