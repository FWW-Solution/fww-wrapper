package adapter

import (
	"bytes"
	"fmt"
	"fww-wrapper/internal/data/dto"
	"fww-wrapper/internal/data/dto_ticket"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/mitchellh/mapstructure"
)

// RedeemTicket implements Adapter.
func (a *adapter) RedeemTicket(body *dto_ticket.Request) (resp dto_ticket.Response, err error) {

	payload, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		return resp, err
	}

	url := fmt.Sprintf("http://%s:%s/api/private/v1/ticket/redeem", a.cfg.Host, a.cfg.Port)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return resp, err
	}

	req.Header.Set("Content-Type", "application/json")

	response, err := a.client.Do(req)
	if err != nil {
		return resp, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return resp, err
	}

	var responseBase dto.BaseResponse

	dec := json.NewDecoder(response.Body)
	if err = dec.Decode(&responseBase); err != nil {
		return resp, err
	}

	if err = mapstructure.Decode(responseBase.Data, &resp); err != nil {
		return resp, err
	}

	return resp, nil

}
