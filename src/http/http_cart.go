package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	atDomain "github.com/mohsanabbas/cart-microservice/src/domain/cart"
	"github.com/mohsanabbas/cart-microservice/src/services/cart"
	"github.com/mohsanabbas/cart-microservice/src/util/rest_errors"
)

// CartHandler http handler interface
type CartHandler interface {
	Create(*gin.Context)
	GetById(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	DeleteAll(*gin.Context)
}

type cartHandler struct {
	service cart.Service
}

// NewCartHandler construct http handlers
func NewCartHandler(service cart.Service) CartHandler {
	return &cartHandler{
		service: service,
	}
}

// Create cart handler
func (handler *cartHandler) Create(c *gin.Context) {
	rh := atDomain.RequestHeaders{}
	if err := c.ShouldBindHeader(&rh); err != nil {
		restErr := rest_errors.NewBadRequestError("Invalid request headers")
		c.JSON(restErr.Status(), restErr)
		return
	}

	request := atDomain.Cart{}

	if err := c.BindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	cart, err := handler.service.Create(request, rh)

	if err != nil {
		restErr := rest_errors.NewInternalServerError("Internal server error", err)
		c.JSON(restErr.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, cart)
}

// GetById handler
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
		restErr := rest_errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	cart, err := handler.service.Update(c.Param("id"), request)
	if err != nil {
		restErr := rest_errors.NewInternalServerError("Internal server error", err)
		c.JSON(restErr.Status(), err)
		return
	}

	c.JSON(http.StatusOK, cart)
}

// Delete handler
func (handler *cartHandler) Delete(c *gin.Context) {
	cart, err := handler.service.Delete(c.Param("cartId"), c.Param("itemId"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, cart)
}

func (handler *cartHandler) DeleteAll(c *gin.Context) {
	cart, err := handler.service.DeleteAll(c.Param("cartId"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, cart)
}
