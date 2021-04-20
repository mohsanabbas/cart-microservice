package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mohsanabbas/cart-microservice/src/http"
	"github.com/mohsanabbas/cart-microservice/src/repository/db"
	"github.com/mohsanabbas/cart-microservice/src/services/cart"
	"github.com/mohsanabbas/ticketing_utils-go/logger"
)
var (
	router = gin.Default()
)

func StartApplication() {
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	atHandler := http.NewCartHandler(cart.NewService(db.NewCartRepository()))

  // App endpoints
	router.POST("/cart", atHandler.Create)


	if err := router.Run(":8080"); err != nil {
		logger.Error(fmt.Sprintf("server has refused to start: %v", nil), err)
	}
}
