package handlers

import (
	"main/component/api"
	"main/component/models"
	"main/component/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ShopHandler struct {
	shopService *services.ShopService
}

func NewShopHandler(shopService *services.ShopService) *ShopHandler {
	return &ShopHandler{
		shopService: shopService,
	}
}

// no need godoc
func (c *ShopHandler) GetData(ctx *gin.Context) {
	// sOffset := ctx.DefaultQuery("offset", "0")
	// sLimit := ctx.DefaultQuery("limit", "10")
	// offset, err := strconv.Atoi(sOffset)
	// if err != nil {
	// 	api.ErrorWithMessage(ctx, http.StatusBadGateway, err.Error())
	// 	return
	// }

	// limit, err := strconv.Atoi(sLimit)
	// if err != nil {
	// 	api.ErrorWithMessage(ctx, http.StatusBadGateway, err.Error())
	// 	return
	// }

	respData, err := c.shopService.GetAllData(ctx)
	if err != nil {
		api.ErrorWithMessage(ctx, http.StatusBadGateway, err.Error())
		return
	}
	api.Ok(ctx, respData)
}

// get by id
func (c *ShopHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	shopId, err := strconv.Atoi(id)
	if err != nil {
		api.ErrorWithMessage(ctx, http.StatusBadGateway, err.Error())
		return
	}
	respData, err := c.shopService.GetById(ctx, shopId)
	if err != nil {
		api.ErrorWithMessage(ctx, http.StatusBadGateway, err.Error())
		return
	}
	api.Ok(ctx, respData)
}

// InsertShop
func (c *ShopHandler) InsertShop(ctx *gin.Context) {
	var data models.Shop
	ctx.BindJSON(&data)
	shop := c.shopService.Insert(ctx, data)
	api.Created(ctx, shop)
}

// UpdateShop
func (c *ShopHandler) UpdateShop(ctx *gin.Context) {
	var data models.Shop
	ctx.BindJSON(&data)

	err := c.shopService.Update(ctx, data)
	if err != nil {
		api.ErrorWithMessage(ctx, http.StatusBadGateway, err.Error())
		return
	}
	api.Ok(ctx, nil)
}
