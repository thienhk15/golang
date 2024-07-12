package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func getMessage(code int) string {
	return http.StatusText(code)
}

func Ok(ctx *gin.Context, data interface{}) {
	response := ApiResponse{
		Code:    http.StatusOK,
		Message: getMessage(http.StatusOK),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, response)
}

func Created(ctx *gin.Context, data interface{}) {
	response := ApiResponse{
		Code:    http.StatusCreated,
		Message: getMessage(http.StatusCreated),
		Data:    data,
	}
	ctx.JSON(http.StatusCreated, response)
}

func Custom(ctx *gin.Context, code int, message string, data interface{}) {
	response := ApiResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
	ctx.JSON(code, response)
}

func Error(ctx *gin.Context, code int) {
	response := ApiResponse{
		Code:    code,
		Message: getMessage(code),
	}
	ctx.JSON(code, response)
}

func ErrorWithMessage(ctx *gin.Context, code int, message string) {
	response := ApiResponse{
		Code:    code,
		Message: message,
	}
	ctx.JSON(code, response)
}

func ErrorWithData(ctx *gin.Context, code int, message string, data interface{}) {
	response := ApiResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
	ctx.JSON(code, response)
}

func ErrorWithRootCause(ctx *gin.Context, code int, message string, errors interface{}) {
	response := ApiResponse{
		Code:    code,
		Message: message,
		Errors:  errors,
	}
	ctx.JSON(code, response)
}
