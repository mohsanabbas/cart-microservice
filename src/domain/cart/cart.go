package cart

import (
	"time"

	"github.com/mohsanabbas/ticketing_utils-go/rest_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	expirationTime = 24
)

type Cart struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Expire       int64              `json:"expire,omitempty" bson:"expire"`
	Items        []Item             `json:"items,omitempty" bson:"items"`
	Name         string             `json:"name,omitempty" bson:"name" `
	AgentSign    string             `json:"agentSign,omitempty" bson:"agentSign"`
	User         string             `json:"user,omitempty" bson:"user"`
	BusinessUnit string             `json:"businessUnit,omitempty" bson:"businessUnit"`
}

type Item struct {
	ID   primitive.ObjectID     `json:"id,omitempty" bson:"_id,omitempty"`
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type CartUpdate struct {
	ModifiedCount int64 `json:"modifiedCount"`
	Results       Cart  `json:"results"`
}

func (at *Cart) Validate() rest_errors.RestErr {

	if at.Expire <= 0 {
		return rest_errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func (at *Cart) SetCartExpiration() {
	at.Expire = time.Now().UTC().Add(expirationTime * time.Hour).Unix()

}
func (id *Item) GenerateItemID() {
	id.ID = primitive.NewObjectID()
}

func (at Cart) IsExpired() bool {
	return time.Unix(at.Expire, 0).Before(time.Now().UTC())
}

func (at *Item) Validate() rest_errors.RestErr {
	if at.Type == "" {
		return rest_errors.NewBadRequestError("invalid product type")
	}
	if at.Data == nil {
		return rest_errors.NewBadRequestError("product data can not be nil")
	}

	return nil
}
