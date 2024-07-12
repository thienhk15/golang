package handlers

import (
	"main/component/api"
	"main/component/models"
	"main/component/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService *services.ProductService
}

func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (c *ProductHandler) GetProducts(ctx *gin.Context) {
	respData, err := c.productService.GetAllData(ctx)
	if err != nil {
		api.ErrorWithMessage(ctx, http.StatusBadGateway, err.Error())
		return
	}
	api.Ok(ctx, respData)
}

func (c *ProductHandler) GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	productId, err := strconv.Atoi(id)
	if err != nil {
		api.ErrorWithMessage(ctx, http.StatusBadRequest, err.Error())
		return
	}
	product, err := c.productService.GetByID(ctx, productId)
	if err != nil {
		api.ErrorWithMessage(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	api.Ok(ctx, product)
}

func (c *ProductHandler) CreateProduct(ctx *gin.Context) {
	//var req requests.ProductCreateRequest
	// if err := ctx.ShouldBindJSON(&req); err != nil {
	// 	api.ErrorWithMessage(ctx, http.StatusBadRequest, err.Error())
	// 	return
	// }

	var product models.Product

	err := c.productService.Insert(ctx, product)
	if err != nil {
		api.ErrorWithMessage(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	api.Ok(ctx, nil)
}

func (c *ProductHandler) UpdateProduct(ctx *gin.Context) {
	//var req requests.ProductUpdateRequest
	// if err := ctx.ShouldBindJSON(&req); err != nil {
	// 	api.ErrorWithMessage(ctx, http.StatusBadRequest, err.Error())
	// 	return
	// }

	var product models.Product

	err := c.productService.Update(ctx, product)
	if err != nil {
		api.ErrorWithMessage(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	api.Ok(ctx, nil)
}

func (c *ProductHandler) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	productId, err := strconv.Atoi(id)
	if err != nil {
		api.ErrorWithMessage(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = c.productService.Delete(ctx, productId)
	if err != nil {
		api.ErrorWithMessage(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	api.Ok(ctx, nil)
}
