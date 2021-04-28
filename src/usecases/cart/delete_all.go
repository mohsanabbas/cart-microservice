package cart

import (
	"strings"

	"github.com/mohsanabbas/cart-microservice/src/domain/cart"
	"github.com/mohsanabbas/cart-microservice/src/util/rest_errors"
)

// DeleteAll usecase method
func (s *usecase) DeleteAll(cartId string) (*cart.CartUpdate, rest_errors.RestErr) {

	cartId = strings.TrimSpace(cartId)
	if len(cartId) == 0 {
		return nil, rest_errors.NewBadRequestError("invalid cart id")
	}

	cart, err := s.dbRepo.DeleteAll(cartId)
	if err != nil {
		return nil, err
	}
	return cart, nil
}
