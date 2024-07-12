package handlers

import (
	"main/component/api"
	"main/component/models"
	requests "main/component/requests"
	"main/component/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
	userService *services.UserService
	shopService *services.ShopService
}

func NewAuthHandler(authService *services.AuthService, userService *services.UserService,
	shopService *services.ShopService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		userService: userService,
		shopService: shopService,
	}
}

func (c *AuthHandler) UserLogin(ctx *gin.Context) {
	var req requests.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		api.ErrorWithMessage(ctx, http.StatusBadRequest, err.Error())
		return
	}

	respData, err := c.authService.LoginUser(ctx, req.Email, req.Password)
	if err != nil {
		api.ErrorWithMessage(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	api.Ok(ctx, respData)
}

func (c *AuthHandler) UserRegister(ctx *gin.Context) {
	var req requests.UserRegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		api.ErrorWithMessage(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := models.User{
		Email:    req.Email,
		Password: req.Password,
		FullName: req.FullName,
	}

	err := c.authService.RegisterUser(ctx, user)
	if err != nil {
		api.ErrorWithMessage(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	api.Created(ctx, nil)
}

func (c *AuthHandler) ShopRegister(ctx *gin.Context) {
	var req requests.ShopRegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		api.ErrorWithMessage(ctx, http.StatusBadRequest, err.Error())
		return
	}

	shop := models.Shop{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}

	err := c.authService.RegisterShop(ctx, shop)
	if err != nil {
		api.ErrorWithMessage(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	api.Created(ctx, nil)
}

func (c *AuthHandler) ShopLogin(ctx *gin.Context) {
	var req requests.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		api.ErrorWithMessage(ctx, http.StatusBadRequest, err.Error())
		return
	}

	respData, err := c.authService.LoginShop(ctx, req.Email, req.Password)
	if err != nil {
		api.ErrorWithMessage(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	api.Ok(ctx, respData)
}

// refresh token
func (c *AuthHandler) RefreshToken(ctx *gin.Context) {
	var req requests.RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		api.ErrorWithMessage(ctx, http.StatusBadRequest, err.Error())
		return
	}

	respData, err := c.authService.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		api.ErrorWithMessage(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	api.Ok(ctx, respData)
}
