package adapter

import (
	"fmt"
	"fww-wrapper/internal/data/dto"
	"fww-wrapper/internal/data/dto_flight"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/mitchellh/mapstructure"
)

// GetFlights implements Adapter.
func (a *adapter) GetFlights(departureTime string, arrivalTime string, limit int, offset int) (resp []dto_flight.ResponseFlight, err error) {
	url := fmt.Sprintf("http://%s:%s/api/private/v1/flights?departure_time=%s&arrival_time=%s&limit%d&offset=%d", a.cfg.Host, a.cfg.Port, departureTime, arrivalTime, limit, offset)

	response, err := a.client.Get(url)
	if err != nil {
		return []dto_flight.ResponseFlight{}, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return []dto_flight.ResponseFlight{}, fmt.Errorf("error status code: %d", response.StatusCode)
	}

	var responseBase dto.BaseResponse

	dec := json.NewDecoder(response.Body)
	if err = dec.Decode(&responseBase); err != nil {
		return
	}

	if err = mapstructure.Decode(responseBase.Data, &resp); err != nil {
		return []dto_flight.ResponseFlight{}, err
	}

	return resp, nil
}

// GetDetailFlightByID implements Adapter.
func (a *adapter) GetDetailFlightByID(id int64) (resp dto_flight.ResponseFlightDetail, err error) {
	url := fmt.Sprintf("http://%s:%s/api/private/v1/flight?id=%d", a.cfg.Host, a.cfg.Port, id)

	response, err := a.client.Get(url)
	if err != nil {
		return dto_flight.ResponseFlightDetail{}, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return dto_flight.ResponseFlightDetail{}, fmt.Errorf("error status code: %d", response.StatusCode)
	}

	var responseBase dto.BaseResponse

	dec := json.NewDecoder(response.Body)
	if err = dec.Decode(&responseBase); err != nil {
		return
	}

	if err = mapstructure.Decode(responseBase.Data, &resp); err != nil {
		return dto_flight.ResponseFlightDetail{}, err
	}

	return resp, nil
}
