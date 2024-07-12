package services

import (
	"context"
	"main/component/models"
	"main/component/repositories"
)

type OrderService struct {
	orderRepo *repositories.OrderRepo
}

func NewOrderService(orderRepo *repositories.OrderRepo) *OrderService {
	os := &OrderService{
		orderRepo: orderRepo,
	}

	// Register new kafka listener
	//os.ListenerFoo1()

	return os
}

func (s *OrderService) GetAllData(ctx context.Context) ([]models.Order, error) {
	listOrders, err := s.orderRepo.GetAll(ctx)
	return listOrders, err
}

func (s *OrderService) Insert(ctx context.Context, data models.Order) error {
	_, err := s.orderRepo.Insert(ctx, data)

	return err
}

func (s *OrderService) Update(ctx context.Context, data models.Order) error {
	_, err := s.orderRepo.Update(ctx, data)
	return err
}

func (s *OrderService) Cancel(ctx context.Context, orderID int) error {
	_, err := s.orderRepo.Cancel(ctx, orderID)
	return err
}

func (s *OrderService) ListenerFoo1() {
	// topic := utils.AppConfig.Kafka.Consumer.FooTopic.Name
	// s.kafkaInstance.Subscribe(topic, func(b []byte) {
	// 	strValue := string(b)
	// 	utils.ShowInfoLogs(fmt.Sprintf("ListenerFoo1 >> %s", strValue))
	// })
}
