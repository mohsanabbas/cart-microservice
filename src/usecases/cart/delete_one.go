package cart

import (
	"strings"

	"github.com/mohsanabbas/cart-microservice/src/domain/cart"
	"github.com/mohsanabbas/cart-microservice/src/util/rest_errors"
)

// Delete usecase method
func (s *usecase) Delete(cartId string, itemId string) (*cart.CartUpdate, rest_errors.RestErr) {

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
