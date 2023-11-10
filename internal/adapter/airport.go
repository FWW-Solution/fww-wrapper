package adapter

import (
	"fmt"
	"fww-wrapper/internal/data/dto"
	"fww-wrapper/internal/data/dto_airport"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/mitchellh/mapstructure"
)

// GetAirport implements Adapter.
func (a *adapter) GetAirport(city string, province string, iata string) (resp []dto_airport.ResponseAirport, err error) {

	url := fmt.Sprintf("http://%s:%s/api/private/v1/airports?city=%s&province=%s&iata=%s", a.cfg.Host, a.cfg.Port, city, province, iata)

	response, err := a.client.Get(url)
	if err != nil {
		return []dto_airport.ResponseAirport{}, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return []dto_airport.ResponseAirport{}, fmt.Errorf("error status code: %d", response.StatusCode)
	}

	var responseBase dto.BaseResponse

	dec := json.NewDecoder(response.Body)
	if err = dec.Decode(&responseBase); err != nil {
		return
	}

	if err = mapstructure.Decode(responseBase.Data, &resp); err != nil {
		return []dto_airport.ResponseAirport{}, err
	}

	return resp, nil
}
