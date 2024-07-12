package services

import (
	"context"
	"main/component/models"
	"main/component/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepo
}

func NewUserService(userRepo *repositories.UserRepo) *UserService {
	us := &UserService{
		userRepo: userRepo,
	}

	// Register new kafka listener
	//us.ListenerFoo1()

	return us
}

func (s *UserService) GetAllData(ctx context.Context) ([]models.User, error) {
	listUsers, err := s.userRepo.GetAll(ctx)
	return listUsers, err
}

func (s *UserService) GetById(ctx context.Context, userId int) (models.User, error) {
	user, err := s.userRepo.GetById(ctx, userId)
	return user, err
}

func (s *UserService) Insert(ctx context.Context, data models.User) error {
	_, err := s.userRepo.Insert(ctx, data)

	return err
}

func (s *UserService) Update(ctx context.Context, data models.User) error {
	_, err := s.userRepo.Update(ctx, data)
	return err
}

func (s *UserService) EmailExists(ctx context.Context, email string) (bool, error) {
	return s.userRepo.EmailExists(ctx, email)
}
func (s *UserService) GeneratePassword(ctx context.Context) (string, error) {
	password, err := s.userRepo.GeneratePassword(ctx)
	if err != nil {
		return "", err
	}
	return password, nil
}

func (s *UserService) ListenerFoo1() {
	// topic := utils.AppConfig.Kafka.Consumer.FooTopic.Name
	// s.kafkaInstance.Subscribe(topic, func(b []byte) {
	// 	strValue := string(b)
	// 	utils.ShowInfoLogs(fmt.Sprintf("ListenerFoo1 >> %s", strValue))
	// })
}
