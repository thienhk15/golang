package services

import (
	"context"
	"main/component/models"
	"main/component/repositories"
)

type ShopService struct {
	shopRepo *repositories.ShopRepo
}

func NewShopService(shopRepo *repositories.ShopRepo) *ShopService {
	ss := &ShopService{
		shopRepo: shopRepo,
	}

	// Register new kafka listener
	//ps.ListenerFoo1()

	return ss
}

func (s *ShopService) GetAllData(ctx context.Context) ([]models.Shop, error) {
	listPosts, err := s.shopRepo.GetAll(ctx)
	return listPosts, err
}

func (s *ShopService) GetById(ctx context.Context, postID int) (models.Shop, error) {
	post, err := s.shopRepo.GetById(ctx, postID)
	return post, err
}

func (s *ShopService) Insert(ctx context.Context, data models.Shop) error {
	_, err := s.shopRepo.Insert(ctx, data)

	return err
}

func (s *ShopService) Update(ctx context.Context, data models.Shop) error {
	_, err := s.shopRepo.Update(ctx, data)
	return err
}

func (s *ShopService) Delete(ctx context.Context, postID int) error {
	err := s.shopRepo.Delete(ctx, postID)
	return err
}

func (s *ShopService) ListenerFoo1() {
	// topic := utils.AppConfig.Kafka.Consumer.FooTopic.Name
	// s.kafkaInstance.Subscribe(topic, func(b []byte) {
	// 	strValue := string(b)
	// 	utils.ShowInfoLogs(fmt.Sprintf("ListenerFoo1 >> %s", strValue))
	// })
}
