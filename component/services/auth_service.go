package services

import (
	"context"
	"errors"
	"main/component/models"
	"main/component/repositories"
	"main/component/responses"
	"time"
)

type AuthService struct {
	userRepo         *repositories.UserRepo
	shopRepo         *repositories.ShopRepo
	refreshTokenRepo *repositories.RefreshTokenRepo
}

func NewAuthService(userRepo *repositories.UserRepo, shopRepo *repositories.ShopRepo, refreshTokenRepo *repositories.RefreshTokenRepo) *AuthService {
	as := &AuthService{
		userRepo:         userRepo,
		shopRepo:         shopRepo,
		refreshTokenRepo: refreshTokenRepo,
	}

	// Register new kafka listener
	//as.ListenerFoo1()

	return as
}

// register user
func (s *AuthService) RegisterUser(ctx context.Context, user models.User) error {
	// Check if email already exists
	user, err := s.userRepo.GetByEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	if user.Id != 0 {
		return errors.New("email already exists")
	}

	// Insert new user
	_, err = s.userRepo.Insert(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// register shop
func (s *AuthService) RegisterShop(ctx context.Context, shop models.Shop) error {
	// Check if shop emails already exists
	shop, err := s.shopRepo.GetByEmail(ctx, shop.Email)
	if err != nil {
		return err
	}
	if shop.Id != 0 {
		return errors.New("email already exists")
	}

	// Insert new shop
	_, err = s.shopRepo.Insert(ctx, shop)
	if err != nil {
		return err
	}

	return nil
}

// login user
func (s *AuthService) LoginUser(ctx context.Context, email string, password string) (responses.UserLoginResponse, error) {
	// Check if email exists
	var res responses.UserLoginResponse
	user, err := s.userRepo.GetByEmailAndPassword(ctx, email, password)
	if err != nil {
		return res, err
	}
	if user.Id == 0 {
		return res, errors.New("email not found")
	}

	// Check password
	if user.Password != password {
		return res, errors.New("wrong password")
	}

	res.User = user
	res.AccessToken = "access_token"
	res.RefreshToken = "refresh"

	return res, nil
}

// login shop
func (s *AuthService) LoginShop(ctx context.Context, email string, password string) (responses.ShopLoginResponse, error) {
	var res responses.ShopLoginResponse
	// Check if email exists
	shop, err := s.shopRepo.GetByEmailAndPassword(ctx, email, password)
	if err != nil {
		return res, err
	}
	if shop.Id == 0 {
		return res, errors.New("email not found")
	}

	// Check password
	if shop.Password != password {
		return res, errors.New("wrong password")
	}

	res.Shop = shop
	res.AccessToken = "access_token"
	res.RefreshToken = "refresh"

	return res, nil
}

// refresh token
func (s *AuthService) RefreshToken(ctx context.Context, token string) (responses.UserLoginResponse, error) {
	var res responses.UserLoginResponse
	// Check if refresh token exists
	refreshToken, err := s.refreshTokenRepo.GetByToken(ctx, token)
	if err != nil {
		return res, err
	}
	if refreshToken.Id == 0 {
		return res, errors.New("refresh token not found")
	}

	// Check if refresh token is expired
	currentTime := time.Now()
	if currentTime.After(refreshToken.ExpiredAt) {
		return res, errors.New("refresh token is expired")
	}
	// create new access token and refresh token
	// TODO: implement jwt

	return res, nil
}
