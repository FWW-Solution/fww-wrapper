package controller

import (
	"fww-wrapper/internal/adapter"
	"fww-wrapper/internal/middleware"

	"go.uber.org/zap"
)

type Controller struct {
	Adapter    adapter.Adapter
	Log        *zap.SugaredLogger
	Middleware middleware.Middleware
}
