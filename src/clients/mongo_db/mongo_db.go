package mongo_db

import (
	"context"
	"fmt"

	"github.com/mohsanabbas/cart-microservice/src/util/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context, conf config.Config) *mongo.Client {
	uri := fmt.Sprintf("%v://%v:%v@%v:27017/shopping-cart",
		conf.DBDriver,
		conf.DBUser,
		conf.DBPass,
		conf.DBSource)
	connection := options.Client().ApplyURI(uri)
	session, err := mongo.Connect(ctx, connection)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf(
		"Successfully connected to database: %v",
		conf.DBName))
	return session
}
