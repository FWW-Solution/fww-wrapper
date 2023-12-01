package controller

import (
	"fww-wrapper/internal/adapter"
	"fww-wrapper/internal/usecase"

	"go.uber.org/zap"
)

type Controller struct {
	Adapter adapter.Adapter
	Log     *zap.SugaredLogger
	UseCase usecase.Usecase
}
