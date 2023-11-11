package adapter

import (
	"bytes"
	"fmt"
	"fww-wrapper/internal/config"
	"fww-wrapper/internal/data/dto"
	"fww-wrapper/internal/data/dto_airport"
	"fww-wrapper/internal/data/dto_flight"
	"fww-wrapper/internal/data/dto_passanger"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/mitchellh/mapstructure"

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
	GetAirport(city, province, iata string) (resp []dto_airport.ResponseAirport, err error)
	GetFlights(departureTime, arrivalTime string, limit int, offset int) (resp []dto_flight.ResponseFlight, err error)
	GetDetailFlightByID(id int64) (resp dto_flight.ResponseFlightDetail, err error)
}

func New(client *circuit.HTTPClient, cfg *config.HttpClientConfig) Adapter {
	return &adapter{
		client: client,
		cfg:    cfg,
	}
}

func (a *adapter) GetPassanger(id int) (resp dto_passanger.ResponseDetail, err error) {
	url := fmt.Sprintf("http://%s:%s/api/private/v1/passanger?id=%d", a.cfg.Host, a.cfg.Port, id)
	response, err := a.client.Get(url)
	if err != nil {
		return dto_passanger.ResponseDetail{}, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return dto_passanger.ResponseDetail{}, fmt.Errorf("error status code: %d", response.StatusCode)
	}

	var responseBase dto.BaseResponse

	dec := json.NewDecoder(response.Body)
	if err = dec.Decode(&responseBase); err != nil {
		return
	}

	if err = mapstructure.Decode(responseBase.Data, &resp); err != nil {
		return dto_passanger.ResponseDetail{}, err
	}

	return resp, nil
}

func (a *adapter) RegisterPassanger(body *dto_passanger.RequestRegister) (resp dto_passanger.ResponseRegistered, err error) {
	payload, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		return dto_passanger.ResponseRegistered{}, err
	}

	url := fmt.Sprintf("http://%s:%s/api/private/v1/passanger", a.cfg.Host, a.cfg.Port)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	response, err := a.client.Do(req)
	if err != nil {
		return dto_passanger.ResponseRegistered{}, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		return dto_passanger.ResponseRegistered{}, fmt.Errorf("error status code: %d", response.StatusCode)
	}

	var responseBase dto.BaseResponse

	dec := json.NewDecoder(response.Body)
	if err = dec.Decode(&responseBase); err != nil {
		return dto_passanger.ResponseRegistered{}, err
	}

	if err = mapstructure.Decode(responseBase.Data, &resp); err != nil {
		return dto_passanger.ResponseRegistered{}, err
	}

	return resp, nil
}

func (a *adapter) UpdatePassanger(body *dto_passanger.RequestUpdate) (resp dto_passanger.ResponseUpdate, err error) {
	payload, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		return dto_passanger.ResponseUpdate{}, err
	}

	url := fmt.Sprintf("http://%s:%s/api/private/v1/passanger", a.cfg.Host, a.cfg.Port)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	response, err := a.client.Do(req)
	if err != nil {
		return dto_passanger.ResponseUpdate{}, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		return dto_passanger.ResponseUpdate{}, fmt.Errorf("error status code: %d", response.StatusCode)
	}

	var responseBase dto.BaseResponse

	dec := json.NewDecoder(response.Body)
	if err = dec.Decode(&responseBase); err != nil {
		return
	}

	fmt.Println(responseBase.Data)

	if err = mapstructure.Decode(responseBase.Data, &resp); err != nil {
		return dto_passanger.ResponseUpdate{}, err
	}

	return resp, nil
}
