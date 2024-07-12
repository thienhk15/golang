package server

import (
	"fmt"
	"main/component/handlers"
	"main/component/repositories"
	"main/component/services"
	"main/internal"
	"main/middleware"
	"main/utils"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

var (
	configPath = "./config/application.yml"
)

type apiServer struct {
	apiMiddleware  *middleware.ApiMiddleware
	authHandler    *handlers.AuthHandler
	userHandler    *handlers.UserHandler
	shopHandler    *handlers.ShopHandler
	orderHandler   *handlers.OrderHandler
	cartHandler    *handlers.CartHandler
	productHandler *handlers.ProductHandler
}

func initConfig() (*utils.Config, error) {
	config := &utils.Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func Initialize() *apiServer {
	var err error

	// config
	utils.AppConfig, err = initConfig()
	if err != nil {
		panic(err)
	}

	// init database
	dbInstance := internal.NewPgDatabase()
	db := dbInstance.ConnectPgDatabase(utils.AppConfig.Database)

	// init Kafka
	// producerTopics := []utils.KafkaTopic{utils.AppConfig.Kafka.Producer.FooTopic}
	// kafkaInstance := internal.NewKafkaInstance(producerTopics)

	// repos
	userRepo := repositories.NewUserRepo(db)
	shopRepo := repositories.NewShopRepo(db)
	refreshTokenRepo := repositories.NewRefreshTokenRepo(db)
	productRepo := repositories.NewProductRepo(db)
	orderRepo := repositories.NewOrderRepo(db)
	cartRepo := repositories.NewCartRepo(db)

	// middleware
	apiMiddleware := middleware.NewMiddleware(userRepo, shopRepo, refreshTokenRepo)

	// services
	userSvc := services.NewUserService(userRepo)
	shopSvc := services.NewShopService(shopRepo)
	authSvc := services.NewAuthService(userRepo, shopRepo, refreshTokenRepo)
	productSvc := services.NewProductService(productRepo)
	orderSvc := services.NewOrderService(orderRepo)
	cartSvc := services.NewCartService(cartRepo)

	// handlers
	userHandler := handlers.NewUserHandler(userSvc)
	shopHandler := handlers.NewShopHandler(shopSvc)
	authHandler := handlers.NewAuthHandler(authSvc, userSvc, shopSvc)
	productHandler := handlers.NewProductHandler(productSvc)
	orderHandler := handlers.NewOrderHandler(orderSvc)
	cartHandler := handlers.NewCartHandler(cartSvc)

	return &apiServer{
		apiMiddleware:  apiMiddleware,
		userHandler:    userHandler,
		shopHandler:    shopHandler,
		authHandler:    authHandler,
		orderHandler:   orderHandler,
		cartHandler:    cartHandler,
		productHandler: productHandler,
	}
}

func (apiServer *apiServer) Start() {
	appConfig := utils.AppConfig
	apiServer.setMode(appConfig.Server.Mode)

	app := gin.New()
	app.SetTrustedProxies(nil)
	app.Use(gin.Logger())
	app.Use(cors.Default())
	// app.Use(...) //middleware

	apiServer.initApiRoutes(app)

	app.Run(fmt.Sprintf(":%d", appConfig.Server.Port))
}

func (apiServer *apiServer) setMode(mode string) {
	switch mode {
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	}
}
