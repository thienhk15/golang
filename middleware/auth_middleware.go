package middleware

import (
	"github.com/gin-gonic/gin"
)

func (mid *ApiMiddleware) AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
