package app

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mohsanabbas/cart-microservice/src/clients/mongo"
	"github.com/mohsanabbas/cart-microservice/src/http"
	"github.com/mohsanabbas/cart-microservice/src/repository/db"
	"github.com/mohsanabbas/cart-microservice/src/services/cart"
	"github.com/mohsanabbas/cart-microservice/src/util/config"
	"github.com/mohsanabbas/ticketing_utils-go/logger"
)

var (
	router = gin.Default()
)

// StartApplication app entry point
func StartApplication() {
	// Get app configs
	config, err := config.GetConfig(".")
	if err != nil {
		logger.Error("connot load config:", err)
	}
	// Create context
	ctx := context.Background()

	// GetSession gets mongo client
	session := mongo.GetSession()
	defer session.Disconnect(ctx)
	client := session.Database(config.DBName).Collection(config.Collection)

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// Logger middleware
	router.Use(gin.Logger())

	// NewCartHandler create handler
	atHandler := http.NewCartHandler(
		cart.NewService(db.NewCartRepository(ctx, client)))

	// App endpoints
	router.POST("/cart", atHandler.Create)
	router.GET("/cart/:id", atHandler.GetById)
	router.PATCH("/cart/:id", atHandler.Update)
	router.DELETE("/cart/:cartId/:itemId", atHandler.Delete)

	if err := router.Run(config.ServerAddress); err != nil {
		logger.Error(fmt.Sprintf("server has refused to start: %v", nil), err)
	}
}
