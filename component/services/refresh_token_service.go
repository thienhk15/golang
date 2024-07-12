package services

import (
	"context"
	"main/component/models"
	"main/component/repositories"
)

type RefreshTokenService struct {
	refreshTokenRepo *repositories.RefreshTokenRepo
}

func NewRefreshTokenService(refreshTokenRepo *repositories.RefreshTokenRepo) *RefreshTokenService {
	rs := &RefreshTokenService{
		refreshTokenRepo: refreshTokenRepo,
	}

	// Register new kafka listener
	//rs.ListenerFoo1()

	return rs
}

func (s *RefreshTokenService) Insert(ctx context.Context, data models.RefreshToken) error {
	_, err := s.refreshTokenRepo.Insert(ctx, data)

	return err
}

func (s *RefreshTokenService) Update(ctx context.Context, data models.RefreshToken) error {
	_, err := s.refreshTokenRepo.Update(ctx, data)
	return err
}

func (s *RefreshTokenService) ListenerFoo1() {
	// topic := utils.AppConfig.Kafka.Consumer.FooTopic.Name
	// s.kafkaInstance.Subscribe(topic, func(b []byte) {
	// 	strValue := string(b)
	// 	utils.ShowInfoLogs(fmt.Sprintf("ListenerFoo1 >> %s", strValue))
	// })
}
