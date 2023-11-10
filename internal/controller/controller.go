package controller

import (
	"fww-wrapper/internal/adapter"

	"go.uber.org/zap"
)

type Controller struct {
	Adapter adapter.Adapter
	Log     *zap.SugaredLogger
}
