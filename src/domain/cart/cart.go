package cart

import (
	"time"

	"github.com/mohsanabbas/cart-microservice/src/util/rest_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	expirationTime = 24
)

// Cart response structure
type Cart struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Expire       int64              `json:"expire" bson:"expire"`
	Items        []Item             `json:"items" bson:"items"`
	BusinessUnit string             `json:"businessUnit" bson:"businessUnit"`
	UserData     User               `json:"userData" bson:"userData"`
}

// User userData structue
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

// Item structure
type Item struct {
	ID   primitive.ObjectID     `json:"id,omitempty" bson:"_id,omitempty"`
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

// CartUpdate response structure after item update in cart
type CartUpdate struct {
	ModifiedCount int64 `json:"modifiedCount"`
	Results       Cart  `json:"results"`
}

// RequestHeaders request headers
type RequestHeaders struct {
	UserToken    string `header:"gtw-sec-user-token" binding:"required"`
	BusinessUnit string `header:"gtw-business-unit" binding:"required"`
}

// ValidateExpiration validate expiry time
func (ct *Cart) ValidateExpiration() rest_errors.RestErr {

	if ct.Expire <= 0 {
		return rest_errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

// SetCartExpiration set expiry time on cart
func (ct *Cart) SetCartExpiration() {
	ct.Expire = time.Now().UTC().Add(expirationTime * time.Hour).Unix()

}

// GenerateItemID adds item "_id" in db
func (it *Item) GenerateItemID() {
	it.ID = primitive.NewObjectID()
}

// IsExpired checks cart expiration time
func (ct Cart) IsExpired() bool {
	return time.Unix(ct.Expire, 0).Before(time.Now().UTC())
}

// Validate checks items request body
func (it *Item) Validate() rest_errors.RestErr {
	if it.Type == "" {
		return rest_errors.NewBadRequestError("invalid product type")
	}
	if it.Data == nil {
		return rest_errors.NewBadRequestError("product data can not be nil")
	}
	return nil
}

// SetUserData adds user data
func (ct *Cart) SetUserData(credential User) {
	ct.UserData = credential
}

// SetBusinessUnit adds business unit
func (ct *Cart) SetBusinessUnit(bu string) {
	ct.BusinessUnit = bu
}
