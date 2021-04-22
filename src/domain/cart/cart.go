package cart

import (
	"time"

	"github.com/mohsanabbas/ticketing_utils-go/rest_errors"
)

const (
	expirationTime = 24
)

type Cart struct {
	ID           interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Expire       int64       `json:"expire,omitempty" bson:"expire"`
	Items        interface{} `json:"items,omitempty" bson:"items"`
	Name         string      `json:"name,omitempty" bson:"name" `
	AgentSign    string      `json:"agentSign,omitempty" bson:"agentSign"`
	User         string      `json:"user,omitempty" bson:"user"`
	BusinessUnit string      `json:"businessUnit,omitempty" bson:"businessUnit"`
}

// type ItemRequest struct {
// 	Type string      `json:"type"`
// 	Data interface{} `json:"data"`
// }

// type Items struct {
// 	Items []interface{} `json:"items,omitempty" bson:"items"`
// }

// type CartResponse struct {
// 	Id    string        `json:"_id"`
// 	Items []interface{} `json:"items"`
// }

func (at *Cart) Validate() rest_errors.RestErr {

	if at.Expire <= 0 {
		return rest_errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func (at *Cart) SetCartExpiration() {
	at.Expire = time.Now().UTC().Add(expirationTime * time.Hour).Unix()

}
func (at Cart) IsExpired() bool {
	return time.Unix(at.Expire, 0).Before(time.Now().UTC())
}
