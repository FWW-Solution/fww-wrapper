package adapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"fww-wrapper/internal/data/dto"
	"fww-wrapper/internal/data/dto_passanger"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

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
	if err != nil {
		return dto_passanger.ResponseRegistered{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	response, err := a.client.Do(req)

	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

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
	if err != nil {
		return dto_passanger.ResponseUpdate{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	response, err := a.client.Do(req)

	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if response.StatusCode != http.StatusCreated {
		return dto_passanger.ResponseUpdate{}, fmt.Errorf("error status code: %d", response.StatusCode)
	}

	var responseBase dto.BaseResponse

	dec := json.NewDecoder(response.Body)
	if err = dec.Decode(&responseBase); err != nil {
		return
	}

	if err = mapstructure.Decode(responseBase.Data, &resp); err != nil {
		return dto_passanger.ResponseUpdate{}, err
	}

	return resp, nil
}
