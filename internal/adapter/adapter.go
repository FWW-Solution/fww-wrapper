package adapter

import (
	"bytes"
	"fmt"
	"fww-wrapper/internal/config"
	"fww-wrapper/internal/data/dto_passanger"
	"net/http"

	"github.com/goccy/go-json"

	circuit "github.com/rubyist/circuitbreaker"
)

type adapter struct {
	client *circuit.HTTPClient
	cfg    *config.HttpClientConfig
}

type Adapter interface {
	GetPassanger(id int) (resp dto_passanger.ResponseDetail, err error)
	RegisterPassanger(body *dto_passanger.RequestRegister) (resp dto_passanger.ResponseRegistered, err error)
	UpdatePassanger(body *dto_passanger.RequestUpdate) (resp dto_passanger.ResponseUpdate, err error)
}

func New(client *circuit.HTTPClient) Adapter {
	return &adapter{
		client: client,
	}
}

func (a *adapter) GetPassanger(id int) (resp dto_passanger.ResponseDetail, err error) {
	url := fmt.Sprintf("%s:%s/private/v1/passanger?id=%d", a.cfg.Host, a.cfg.Port, id)
	response, err := a.client.Get(url)
	if err != nil {
		return dto_passanger.ResponseDetail{}, err
	}

	defer response.Body.Close()

	dec := json.NewDecoder(response.Body)
	if err = dec.Decode(&resp); err != nil {
		return
	}

	return resp, nil
}

func (a *adapter) RegisterPassanger(body *dto_passanger.RequestRegister) (resp dto_passanger.ResponseRegistered, err error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return dto_passanger.ResponseRegistered{}, err
	}

	url := fmt.Sprintf("%s:%s/private/v1/passanger", a.cfg.Host, a.cfg.Port)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	response, err := a.client.Do(req)
	if err != nil {
		return dto_passanger.ResponseRegistered{}, err
	}

	defer response.Body.Close()

	dec := json.NewDecoder(response.Body)
	if err = dec.Decode(&resp); err != nil {
		return
	}

	return resp, nil
}

func (a *adapter) UpdatePassanger(body *dto_passanger.RequestUpdate) (resp dto_passanger.ResponseUpdate, err error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return dto_passanger.ResponseUpdate{}, err
	}

	url := fmt.Sprintf("%s:%s/private/v1/passanger", a.cfg.Host, a.cfg.Port)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	response, err := a.client.Do(req)
	if err != nil {
		return dto_passanger.ResponseUpdate{}, err
	}

	defer response.Body.Close()

	dec := json.NewDecoder(response.Body)
	if err = dec.Decode(&resp); err != nil {
		return
	}

	return resp, nil
}
