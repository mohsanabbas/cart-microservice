package app

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mohsanabbas/cart-microservice/src/clients/mongo_db"
	"github.com/mohsanabbas/cart-microservice/src/http"
	"github.com/mohsanabbas/cart-microservice/src/repository/db"
	"github.com/mohsanabbas/cart-microservice/src/services/cart"
	"github.com/mohsanabbas/cart-microservice/src/util"
	"github.com/mohsanabbas/ticketing_utils-go/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	// Get app configs
	config, err := util.GetConfig(".")
	if err != nil {
		logger.Error("connot load config:", err)
	}
	// Create context
	ctx := context.TODO()

	// Instanciate mongo client
	session := mongo_db.Connect(ctx, config)
	client := session.Database(config.DBName).Collection(config.Collection)

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// Logger middleware
	router.Use(gin.Logger())

	// Create handler
	atHandler := http.NewCartHandler(
		cart.NewService(db.NewCartRepository(client, ctx)))

	// App endpoints
	router.POST("/cart", atHandler.Create)
	router.GET("/cart/:id", atHandler.GetById)

	if err := router.Run(config.ServerAddress); err != nil {
		logger.Error(fmt.Sprintf("server has refused to start: %v", nil), err)
	}
}
