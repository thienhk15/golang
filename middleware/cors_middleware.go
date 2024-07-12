package middleware

import (
	"main/utils"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (m *ApiMiddleware) CorsMiddleware(conf utils.CorsConfig) gin.HandlerFunc {
	if !conf.IsEnabled {
		return func(ctx *gin.Context) {
			ctx.Next()
		}
	}
	return cors.New(cors.Config{
		AllowOrigins:     conf.AllowedOrigins,
		AllowMethods:     conf.AllowedMethods,
		AllowHeaders:     conf.AllowedHeaders,
		ExposeHeaders:    conf.ExposedHeaders,
		AllowCredentials: conf.AllowCredentials,
		MaxAge:           time.Duration(conf.MaxAge),
	})
}
