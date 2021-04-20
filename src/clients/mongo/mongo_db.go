package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Cart struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Expires   string             `bson:"expires,omitempty"`
	Items     []string           `bson:"items,omitempty"`
	AgentSign string             `bson:"agentSign"`
}

func init() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	database := client.Database("ms-shopping-cart")
	cartCollection := database.Collection("cart")

	cart := Cart{
		Expires: "Nic Raboy",
		Items:   []string{"development", "programming", "coding"},
	}
	insertResult, err := cartCollection.InsertOne(ctx, cart)
	if err != nil {
		panic(err)
	}
	fmt.Println(insertResult.InsertedID)
}
