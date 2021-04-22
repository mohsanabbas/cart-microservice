package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	atDomain "github.com/mohsanabbas/cart-microservice/src/domain/cart"
	"github.com/mohsanabbas/cart-microservice/src/services/cart"
	"github.com/mohsanabbas/ticketing_utils-go/rest_errors"
)

type CartHandler interface {
	Create(*gin.Context)
	GetById(*gin.Context)
}

type cartHandler struct {
	service cart.Service
}

func NewCartHandler(service cart.Service) CartHandler {
	return &cartHandler{
		service: service,
	}
}

func (handler *cartHandler) Create(c *gin.Context) {
	request := atDomain.Cart{}

	if err := c.BindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	cart, err := handler.service.Create(request)

	if err != nil {
		restErr := rest_errors.NewInternalServerError("iternal server error", err)
		c.JSON(restErr.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, cart)
}

func (handler *cartHandler) GetById(c *gin.Context) {
	accessToken, err := handler.service.GetById(c.Param("id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}
