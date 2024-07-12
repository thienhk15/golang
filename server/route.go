package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (apiServer *apiServer) initApiRoutes(app *gin.Engine) {
	userApi := app.Group("/api/user", apiServer.apiMiddleware.AuthMiddleWare())
	{
		userApi.GET("", apiServer.userHandler.GetData)
		// apiGroup.POST("", apiServer.userHandler.InsertData)
		userApi.GET("/:id, apiServer.userHandler.GetDataById")
		userApi.PUT("/:id", apiServer.userHandler.UpdateData)
	}

	productApi := app.Group("/api/product", apiServer.apiMiddleware.AuthMiddleWare())
	{
		productApi.GET("", apiServer.productHandler.GetProducts)
		productApi.POST("", apiServer.productHandler.CreateProduct)
		productApi.GET("/:id", apiServer.productHandler.GetProduct)
		productApi.PUT("/:id", apiServer.productHandler.UpdateProduct)
	}

	shopApi := app.Group("/api/shop", apiServer.apiMiddleware.AuthMiddleWare())
	{
		shopApi.GET("", apiServer.shopHandler.GetData)
		shopApi.POST("", apiServer.shopHandler.InsertShop)
		shopApi.GET("/:id", apiServer.shopHandler.GetByID)
		shopApi.PUT("/:id", apiServer.shopHandler.UpdateShop)
	}

	authApi := app.Group("/api/auth")
	{
		authApi.POST("/login/user", apiServer.authHandler.UserLogin)
		authApi.POST("/register/user", apiServer.authHandler.UserRegister)
		authApi.POST("/login/shop", apiServer.authHandler.ShopLogin)
		authApi.POST("/register/shop", apiServer.authHandler.ShopRegister)
		authApi.POST("/refreshToken", apiServer.authHandler.RefreshToken)
	}

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
