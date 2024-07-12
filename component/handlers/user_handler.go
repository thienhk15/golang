package handlers

import (
	"main/component/api"
	"main/component/models"
	"main/component/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetUserData godoc
//
//	@Summary		Get user data
//	@Description	Get list user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]models.User
//	@Router			/users 	[get]
func (c *UserHandler) GetData(ctx *gin.Context) {
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

	respData, err := c.userService.GetAllData(ctx)
	if err != nil {
		api.ErrorWithMessage(ctx, http.StatusBadGateway, err.Error())
		return
	}
	api.Ok(ctx, respData)
}

// InsertUserData godoc
//
//	@Summary		Insert user data
//	@Description	Insert user data
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			offset 	body     	models.User  	true  	"User body"
//	@Success		200	{object}	int
//	@Router			/user 	[post]
// func (c *UserHandler) InsertData(ctx *gin.Context) {
// 	user := models.User{}
// 	err := ctx.ShouldBind(&user)
// 	if err == nil {
// 		err := c.userService.Insert(ctx, user)
// 		if err != nil {
// 			api.ErrorWithMessage(ctx, http.StatusBadGateway, err.Error())
// 			return
// 		}
// 		api.Ok(ctx)
// 	} else {
// 		api.ErrorWithMessage(ctx, http.StatusBadRequest, err.Error())
// 	}
// }

// get one
func (c *UserHandler) GetOne(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		api.ErrorWithMessage(ctx, http.StatusBadGateway, err.Error())
		return
	}
	respData, err := c.userService.GetById(ctx, userId)
	if err != nil {
		api.ErrorWithMessage(ctx, http.StatusBadGateway, err.Error())
		return
	}
	api.Ok(ctx, respData)
}

// func update user
func (c *UserHandler) UpdateData(ctx *gin.Context) {
	user := models.User{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		api.ErrorWithMessage(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = c.userService.Update(ctx, user)
	if err != nil {
		api.ErrorWithMessage(ctx, http.StatusBadGateway, err.Error())
		return
	}

	api.Ok(ctx, nil)
}
