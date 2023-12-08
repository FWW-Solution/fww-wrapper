package adapter

import (
	"bytes"
	"fmt"
	"fww-wrapper/internal/data/dto"
	"fww-wrapper/internal/data/dto_booking"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/mitchellh/mapstructure"
)

// // Booking implements Adapter.
// func (a *adapter) Booking(body *dto_booking.Request) (resp dto_booking.AsyncBookResponse, err error) {
// 	json, err := json.Marshal(body)
// 	if err != nil {
// 		return resp, err
// 	}

// 	ID := watermill.NewUUID()

// 	err = a.publisher.Publish("request_booking", message.NewMessage(
// 		ID,
// 		json,
// 	))
// 	if err != nil {
// 		return resp, err
// 	}

// 	resp = dto_booking.AsyncBookResponse{
// 		BookingIDCode: ID,
// 	}

// 	return resp, nil

// }

func (a *adapter) Booking(body *dto_booking.Request) (resp dto_booking.AsyncBookResponse, err error) {
	url := fmt.Sprintf("http://%s:%s/api/private/v1/booking", a.cfg.Host, a.cfg.Port)

	jsonBody, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		return resp, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
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
		return
	}

	if err = mapstructure.Decode(responseBase.Data, &resp); err != nil {
		return dto_booking.AsyncBookResponse{}, err
	}

	return resp, nil
}

// GetDetailBooking implements Adapter.
func (a *adapter) GetDetailBooking(codeBooking string) (resp dto_booking.BookResponse, err error) {
	url := fmt.Sprintf("http://%s:%s/api/private/v1/booking?code_booking=%s", a.cfg.Host, a.cfg.Port, codeBooking)

	response, err := a.client.Get(url)
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
		return
	}

	if err = mapstructure.Decode(responseBase.Data, &resp); err != nil {
		return dto_booking.BookResponse{}, err
	}

	return resp, nil
}
