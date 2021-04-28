package cart

import (
	"strings"

	"github.com/mohsanabbas/cart-microservice/src/domain/cart"
	"github.com/mohsanabbas/cart-microservice/src/util/rest_errors"
)

// GetById usecase method
func (s *usecase) GetById(id string) (*cart.Cart, rest_errors.RestErr) {
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
