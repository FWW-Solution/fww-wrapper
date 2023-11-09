package main

import (
	"fww-wrapper/internal/config"
	"fww-wrapper/internal/container"
	"fww-wrapper/internal/container/infrastructure/http"
)

func main() {
	// init config
	cfg := config.InitConfig()

	// init service
	app := container.InitService(cfg)

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
