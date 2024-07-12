package middleware

import (
	"main/component/repositories"
	"main/component/services"
)

type ApiMiddleware struct {
	authSvc *services.AuthService
}

func NewMiddleware(userRepo *repositories.UserRepo, shopRepo *repositories.ShopRepo, refreshTokenRepo *repositories.RefreshTokenRepo) *ApiMiddleware {
	return &ApiMiddleware{
		authSvc: services.NewAuthService(userRepo, shopRepo, refreshTokenRepo),
	}
}
