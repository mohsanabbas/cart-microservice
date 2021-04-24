package cart

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/mohsanabbas/cart-microservice/src/domain/cart"
	"github.com/mohsanabbas/cart-microservice/src/repository/db"
	"github.com/mohsanabbas/cart-microservice/src/util/decrypt"
	"github.com/mohsanabbas/ticketing_utils-go/rest_errors"
)

// Service interface
type Service interface {
	Create(cart.Cart, cart.RequestHeaders) (*cart.Cart, rest_errors.RestErr)
	GetById(string) (*cart.Cart, rest_errors.RestErr)
	Update(string, cart.Item) (*cart.CartUpdate, rest_errors.RestErr)
	Delete(string, string) (*cart.CartUpdate, rest_errors.RestErr)
}

type service struct {
	dbRepo db.DbRepository
}

// NewService
func NewService(dbRepo db.DbRepository) Service {
	return &service{
		dbRepo: dbRepo,
	}
}

// Create service method
func (s *service) Create(request cart.Cart, rh cart.RequestHeaders) (*cart.Cart, rest_errors.RestErr) {

	request.SetCartExpiration()
	request.SetBusinessUnit(rh.BusinessUnit)
	credential := cart.User{}
	decoded := decrypt.Decrypt(rh.UserToken)
	if err := json.Unmarshal(decoded, &credential); err != nil {
		return nil,
			rest_errors.NewInternalServerError("Invalid gtw-user-token",
				errors.New("json parsing error"))
	}
	request.SetUserData(credential)

	for k := range request.Items {
		request.Items[k].GenerateItemID()
	}
	cart, err := s.dbRepo.Create(request)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

// GetById service method
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

// Update service method
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
	return cart, nil
}

// Delete service method
func (s *service) Delete(cartId string, itemId string) (*cart.CartUpdate, rest_errors.RestErr) {

	cartId = strings.TrimSpace(cartId)
	if len(cartId) == 0 {
		return nil, rest_errors.NewBadRequestError("invalid cart id")
	}
	itemId = strings.TrimSpace(itemId)
	if len(itemId) == 0 {
		return nil, rest_errors.NewBadRequestError("invalid cart id")
	}
	cart, err := s.dbRepo.Delete(cartId, itemId)
	if err != nil {
		return nil, err
	}
	return cart, nil
}
