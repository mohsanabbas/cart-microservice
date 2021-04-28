package cart

import (
	"strings"

	"github.com/mohsanabbas/cart-microservice/src/domain/cart"
	"github.com/mohsanabbas/cart-microservice/src/util/rest_errors"
)

// Update usecase method
func (s *usecase) Update(id string, request cart.Item) (*cart.CartUpdate, rest_errors.RestErr) {
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
