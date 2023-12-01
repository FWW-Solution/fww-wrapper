package container

import (
	"fww-wrapper/internal/adapter"
	"fww-wrapper/internal/config"
	"fww-wrapper/internal/container/infrastructure/database"
	grpcclient "fww-wrapper/internal/container/infrastructure/grpc/client"
	grpcserver "fww-wrapper/internal/container/infrastructure/grpc/server"
	"fww-wrapper/internal/container/infrastructure/http"
	"fww-wrapper/internal/container/infrastructure/http/router"
	httpclient "fww-wrapper/internal/container/infrastructure/http_client"
	logger "fww-wrapper/internal/container/infrastructure/log"
	messagestream "fww-wrapper/internal/container/infrastructure/message_stream"
	"fww-wrapper/internal/container/infrastructure/redis"
	"fww-wrapper/internal/controller"
	"fww-wrapper/internal/middleware"
	"fww-wrapper/internal/repository"
	"fww-wrapper/internal/usecase"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gofiber/fiber/v2"
)

func InitService(cfg *config.Config) (*fiber.App, []*message.Router) {
	// init database
	db := database.GetConnection(&cfg.Database)
	// init redis
	clientRedis := redis.SetupClient(&cfg.Redis)
	// init redis cache
	redis.InitRedisClient(clientRedis)
	// Init Tracing
	// Init Logger
	log := logger.Initialize(cfg)
	// Init HTTP Server
	server := http.SetupHttpEngine()
	// Init GRPC Server
	grpcserver.Init(&cfg.GrpcServer)
	// Init GRPC Client
	_, err := grpcclient.Init(&cfg.GrpcClient)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	// Init Circuit Breaker
	cb := httpclient.InitCircuitBreaker(&cfg.HttpClient, "consecutive")
	// Init HTTP Client
	client := httpclient.InitHttpClient(&cfg.HttpClient, cb)

	amqpMessageStream := messagestream.NewAmpq(&cfg.MessageStream)

	// set message stream subscriber
	sub, err := amqpMessageStream.NewSubscriber()
	if err != nil {
		log.Error(err)
		panic(err)
	}

	// set message stream publisher
	pub, err := amqpMessageStream.NewPublisher()
	if err != nil {
		log.Error(err)
		panic(err)
	}

	// Init Repository
	repository := repository.NewRepository(db)

	// Init Middleware
	middleware := middleware.Middleware{Repository: repository}

	// Init Adapter
	adapter := adapter.New(client, &cfg.HttpClient, pub, &cfg.Email)
	// Init UseCase
	usecase := usecase.NewUsecase(repository)
	// Init Controller
	ctrl := controller.Controller{Adapter: adapter, Log: log, UseCase: usecase}

	// Init router
	var messageRouters []*message.Router

	sendEmailNotificationRouter, err := messagestream.NewRouter(
		pub,
		"send_email_notification_poisoned",
		"send_email_notification_handler",
		"send_email_notification",
		sub,
		ctrl.SendEmailNotificationHandler,
	)

	if err != nil {
		log.Fatal(err)
	}

	messageRouters = append(messageRouters, sendEmailNotificationRouter)

	// Init Router
	router := router.Initialize(server, &ctrl, &middleware)

	return router, messageRouters
}
