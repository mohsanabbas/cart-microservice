package mongo

import (
	"context"
	"fmt"

	"github.com/mohsanabbas/cart-microservice/src/util/config"
	"github.com/mohsanabbas/ticketing_utils-go/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	session *mongo.Client
)

// init
func init() {

	// Get config
	config, err := config.GetConfig(".")
	if err != nil {
		logger.Error("connot load config:", err)
	}

	// Create context
	ctx := context.TODO()

	// setup uri
	uri := fmt.Sprintf("%v://%v:%v@%v:27017/shopping-cart",
		config.DBDriver,
		config.DBUser,
		config.DBPass,
		config.DBSource)

	// Connect to Mongo:
	connection := options.Client().ApplyURI(uri)
	if session, err = mongo.Connect(ctx, connection); err != nil {
		panic(err)
	}
}

// GetSession return mongo session
func GetSession() *mongo.Client {
	return session
}
