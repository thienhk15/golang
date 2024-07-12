package services

import (
	"context"
	"main/component/models"
	"main/component/repositories"
)

type CartService struct {
	cartRepo *repositories.CartRepo
}

func NewCartService(cartRepo *repositories.CartRepo) *CartService {
	cs := &CartService{
		cartRepo: cartRepo,
	}

	// Register new kafka listener
	//cs.ListenerFoo1()

	return cs
}

func (s *CartService) GetAllData(ctx context.Context) ([]models.Cart, error) {
	listCarts, err := s.cartRepo.GetAll(ctx)
	return listCarts, err
}

func (s *CartService) Insert(ctx context.Context, data models.Cart) error {
	_, err := s.cartRepo.Insert(ctx, data)

	return err
}

func (s *CartService) Update(ctx context.Context, data models.Cart) error {
	_, err := s.cartRepo.Update(ctx, data)
	return err
}

func (s *CartService) Delete(ctx context.Context, cartID int) error {
	err := s.cartRepo.Delete(ctx, cartID)
	return err
}

func (s *CartService) ListenerFoo1() {
	// topic := utils.AppConfig.Kafka.Consumer.FooTopic.Name
	// s.kafkaInstance.Subscribe(topic, func(b []byte) {
	// 	strValue := string(b)
	// 	utils.ShowInfoLogs(fmt.Sprintf("ListenerFoo1 >> %s", strValue))
	// })
}
