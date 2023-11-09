package container

import (
	"fww-wrapper/internal/adapter"
	"fww-wrapper/internal/config"
	grpcclient "fww-wrapper/internal/container/infrastructure/grpc/client"
	grpcserver "fww-wrapper/internal/container/infrastructure/grpc/server"
	"fww-wrapper/internal/container/infrastructure/http"
	"fww-wrapper/internal/container/infrastructure/http/router"
	httpclient "fww-wrapper/internal/container/infrastructure/http_client"
	logger "fww-wrapper/internal/container/infrastructure/log"
	messagestream "fww-wrapper/internal/container/infrastructure/message_stream"
	"fww-wrapper/internal/container/infrastructure/redis"
	"fww-wrapper/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func InitService(cfg *config.Config) *fiber.App {
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
	_, err = amqpMessageStream.NewSubscriber()
	if err != nil {
		log.Error(err)
		panic(err)
	}

	// set message stream publisher
	_, err = amqpMessageStream.NewPublisher()
	if err != nil {
		log.Error(err)
		panic(err)
	}

	// Init Adapter
	adapter := adapter.New(client)
	// Init Controller
	ctrl := controller.Controller{Adapter: adapter, Log: log}

	// Init Router
	router := router.Initialize(server, &ctrl)

	return router
}
