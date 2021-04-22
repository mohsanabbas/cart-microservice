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
	Update(*gin.Context)
	Delete(*gin.Context)
}

type cartHandler struct {
	service cart.Service
}

func NewCartHandler(service cart.Service) CartHandler {
	return &cartHandler{
		service: service,
	}
}

// Create cart handler
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

// Get handler
func (handler *cartHandler) GetById(c *gin.Context) {
	cart, err := handler.service.GetById(c.Param("id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, cart)
}

// Update handler
func (handler *cartHandler) Update(c *gin.Context) {
	request := atDomain.Item{}
	if err := c.BindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	cart, err := handler.service.Update(c.Param("id"), request)

	if err != nil {
		restErr := rest_errors.NewInternalServerError("iternal server error", err)
		c.JSON(restErr.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, cart)
}

// Delete
func (handler *cartHandler) Delete(c *gin.Context) {

	cart, err := handler.service.Delete(c.Param("cartId"), c.Param("itemId"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, cart)
}
