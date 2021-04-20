package cart

import (
	"time"

	"github.com/mohsanabbas/ticketing_utils-go/rest_errors"
)



const (
	expirationTime = 24
)

type Cart struct {
	ID           string        `json:"_id"`
	Expire       int64        `json:"expire,omitempty"`
	Items        []interface{} `json:"items"`
	Name         string        `json:"name,omitempty" `
	AgentSign    string        `json:"agentSign,omitempty"`
	User         string        `json:"user,omitempty"`
	BusinessUnit string        `json:"businessUnit"`
}

type ItemRequest struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type Items struct{
	Items []ItemRequest `json:"items"`
}

type CartResponse struct {
	Id    string        `json:"_id"`
	Items []interface{} `json:"items"`
}

func (at *Cart) Validate() rest_errors.RestErr {

	if at.Expire <= 0 {
		return rest_errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}


func GetNewCart(items []interface{}) Cart {
	return Cart{
		Items:   items,
		Expire: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at Cart) IsExpired() bool {
	return time.Unix(at.Expire, 0).Before(time.Now().UTC())
}
