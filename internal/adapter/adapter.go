package adapter

import (
	"fww-wrapper/internal/data/dto_passanger"

	circuit "github.com/rubyist/circuitbreaker"
)

type adapter struct {
	client *circuit.HTTPClient
}

type Adapter interface {
	GetPassanger(id int) (resp []byte, err error)
	RegisterPassanger(body *dto_passanger.RequestRegister) (resp dto_passanger.ResponseRegistered, err error)
	UpdatePassanger(body *dto_passanger.RequestUpdate) (resp dto_passanger.ResponseUpdate, err error)
}

func New(client *circuit.HTTPClient) Adapter {
	return &adapter{
		client: client,
	}
}

func (a *adapter) GetPassanger(id int) (resp []byte, err error) {
	return nil, nil
}

func (a *adapter) RegisterPassanger(body *dto_passanger.RequestRegister) (resp dto_passanger.ResponseRegistered, err error) {
	return dto_passanger.ResponseRegistered{}, nil
}

func (a *adapter) UpdatePassanger(body *dto_passanger.RequestUpdate) (resp dto_passanger.ResponseUpdate, err error) {
	return dto_passanger.ResponseUpdate{}, nil
}
