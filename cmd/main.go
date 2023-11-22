package main

import (
	"context"
	"fww-wrapper/internal/config"
	"fww-wrapper/internal/container"
	"fww-wrapper/internal/container/infrastructure/http"
	"log"

	"github.com/ThreeDotsLabs/watermill/message"
)

func main() {
	// init config
	cfg := config.InitConfig()

	// init service
	app, routers := container.InitService(cfg)

	for _, router := range routers {
		ctx := context.Background()
		go func(router *message.Router) {
			err := router.Run(ctx)
			if err != nil {
				log.Fatal(err)
			}
		}(router)
	}

	http.StartHttpServer(app, cfg.HttpServer.Port)

	// // init config
	// cfg := config.InitConfig()

	// log := logger.Initialize(cfg)
	// count := 0
	// for {

	// 	if rand.Float32() > 0.8 {
	// 		log.Error("oops...something is wrong",
	// 			zap.Int("count", count),
	// 			zap.Error(errors.New("error details")))
	// 	} else {
	// 		log.Info("everything is fine",
	// 			zap.Int("count", count))
	// 	}
	// 	count++
	// 	time.Sleep(time.Second * 2)
	// }
}
