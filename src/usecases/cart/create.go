package cart

import (
	"encoding/json"
	"errors"

	"github.com/mohsanabbas/cart-microservice/src/domain/cart"
	"github.com/mohsanabbas/cart-microservice/src/util/decrypt"
	"github.com/mohsanabbas/cart-microservice/src/util/rest_errors"
)

// Create usecase method
func (s *usecase) Create(request cart.Cart, rh cart.RequestHeaders) (*cart.Cart, rest_errors.RestErr) {

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
