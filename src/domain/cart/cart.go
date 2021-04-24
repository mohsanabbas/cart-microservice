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
	Expire       int64              `json:"expire" bson:"expire"`
	Items        []Item             `json:"items" bson:"items"`
	BusinessUnit string             `json:"businessUnit" bson:"businessUnit"`
	UserData     User               `json:"userData" bson:"userData"`
}

type User struct {
	Credential struct {
		Personid   int         `json:"personId" bson:"personId"`
		Userid     int         `json:"userId" bson:"userId"`
		VendorName string      `json:"name" bson:"name"`
		Email      string      `json:"email" bson:"email"`
		Cpf        interface{} `json:"cpf" bson:"cpf"`
		Branchid   int         `json:"branchId" bson:"branchId"`
		Agentsign  string      `json:"agentSign" bson:"agentSign"`
		User       string      `json:"user" bson:"user"`
		Usertype   string      `json:"userType" bson:"userType"`
	} `json:"credential" bson:"credential"`
	Systems []interface{} `json:"systems" bson:"systems"`
	Iat     int64         `json:"iat" bson:"iat"`
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

type RequestHeaders struct {
	UserToken    string `json:"gtw-sec-user-token"`
	BusinessUnit string `json:"gtw-business-unit"`
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

// Request headers validation
func (at *RequestHeaders) ValidateHeaders() rest_errors.RestErr {
	if len(at.UserToken) == 0 {
		return rest_errors.NewBadRequestError("gtw-sec-user-token request header is required")
	}
	if at.BusinessUnit == "" {
		return rest_errors.NewBadRequestError("gtw-business-unit request header is required")
	}
	return nil
}

// Set User data
func (ct *Cart) SetUserData(credential User) {
	ct.UserData = credential
}

// Set BusinessUnit
func (ct *Cart) SetBusinessUnit(bu string) {
	ct.BusinessUnit = bu
}
