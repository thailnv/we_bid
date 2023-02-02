package handlers

import (
	"log"
	"net/http"
	"we_bid/models"
	"we_bid/services"

	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	Handler
	getProductsList(context *gin.Context)
	getProductByID(context *gin.Context)
}

type handler struct {
	service services.ProductService
}

func CreateProductHandler(service services.ProductService) ProductHandler {
	return &handler{service: service}
}

func (h *handler) Register(router gin.IRouter) {
	router.GET("/products", h.getProductsList)
	router.GET("/products/:id", h.getProductByID)
}

type GetAllProductsRequest struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

type GetProductsListResponse struct {
	Code     int               `json:"code"`
	Products *[]models.Product `json:"data"`
}

func (h *handler) getProductsList(c *gin.Context) {
	var query GetAllProductsRequest
	err := c.ShouldBindQuery(&query)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, NewInternalServerError(err))
		return
	}
	products, err := h.service.GetProductsList(services.ProductsFilter{Limit: query.Limit, Page: query.Page})
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewInternalServerError(err))
		return
	}
	c.JSON(http.StatusOK, GetProductsListResponse{
		Code:     http.StatusOK,
		Products: products,
	})
}

type GetProductByIDRequest struct {
	ID int `uri:"id" binding:"required"`
}

type GetProductByIDResponse struct {
	Code    int             `json:"code"`
	Product *models.Product `json:"data"`
}

func (h *handler) getProductByID(c *gin.Context) {
	var params GetProductByIDRequest
	err := c.ShouldBindUri(&params)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, NewInternalServerError(err))
		return
	}
	product, err := h.service.GetProductByID(params.ID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, NewRecordNotFoundError(err))
		return
	}
	c.JSON(http.StatusOK, GetProductByIDResponse{Code: http.StatusOK, Product: product})
}
