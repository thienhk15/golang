package utils

import (
	"main/component/responses"

	"github.com/gin-gonic/gin"
)

func GetCurrentUser(ctx *gin.Context) responses.UserLoginResponse {
	data := ctx.Value(ContextUserAuthorized)
	if data != nil {
		return data.(responses.UserLoginResponse)
	}
	return responses.UserLoginResponse{}
}
