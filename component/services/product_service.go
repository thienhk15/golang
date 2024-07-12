package services

import (
	"context"
	"main/component/models"
	"main/component/repositories"
)

type ProductService struct {
	productRepo *repositories.ProductRepo
}

func NewProductService(productRepo *repositories.ProductRepo) *ProductService {
	ps := &ProductService{
		productRepo: productRepo,
	}

	// Register new kafka listener
	//ps.ListenerFoo1()

	return ps
}

func (s *ProductService) GetAllData(ctx context.Context) ([]models.Product, error) {
	listProducts, err := s.productRepo.GetAll(ctx)
	return listProducts, err
}

func (s *ProductService) GetByID(ctx context.Context, productID int) (models.Product, error) {
	product, err := s.productRepo.GetByID(ctx, productID)
	return product, err
}

func (s *ProductService) Insert(ctx context.Context, data models.Product) error {
	_, err := s.productRepo.Insert(ctx, data)

	return err
}

func (s *ProductService) Update(ctx context.Context, data models.Product) error {
	_, err := s.productRepo.Update(ctx, data)
	return err
}

func (s *ProductService) Delete(ctx context.Context, productID int) error {
	err := s.productRepo.Delete(ctx, productID)
	return err
}

func (s *ProductService) ListenerFoo1() {
	// topic := utils.AppConfig.Kafka.Consumer.FooTopic.Name
	// s.kafkaInstance.Subscribe(topic, func(b []byte) {
	// 	strValue := string(b)
	// 	utils.ShowInfoLogs(fmt.Sprintf("ListenerFoo1 >> %s", strValue))
	// })
}
