package adapter

import (
	"encoding/json"
	"errors"
	"fmt"
	"fww-wrapper/internal/data/dto"
	"fww-wrapper/internal/data/dto_payment"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/mitchellh/mapstructure"
)

// DoPayment implements Adapter.
func (a *adapter) DoPayment(body *dto_payment.Request) (resp dto_payment.AsyncPaymentResponse, err error) {
	json, err := json.Marshal(body)
	if err != nil {
		return resp, err
	}

	ID := watermill.NewUUID()

	err = a.publisher.Publish("do_payment", message.NewMessage(
		ID,
		json,
	))
	if err != nil {
		return resp, err
	}

	resp = dto_payment.AsyncPaymentResponse{
		PaymentIDCode: ID,
	}

	return resp, nil
}

// GetPaymentStatus implements Adapter.
func (a *adapter) GetPaymentStatus(paymentCode string) (resp dto_payment.StatusResponse, err error) {
	url := fmt.Sprintf("http://%s:%s/api/private/v1/payment/status?payment_code=%s", a.cfg.Host, a.cfg.Port, paymentCode)

	response, err := a.client.Get(url)
	if err != nil {
		return resp, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		errorMsg := fmt.Sprintf("error get payment methods, status code: %d", response.StatusCode)
		return resp, errors.New(errorMsg)
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

// GetPaymentMethods implements Adapter.
func (a *adapter) GetPaymentMethods() (resp []dto_payment.MethodResponse, err error) {
	url := fmt.Sprintf("http://%s:%s/api/private/v1/payment/methods", a.cfg.Host, a.cfg.Port)

	response, err := a.client.Get(url)
	if err != nil {
		return resp, err
	}

	defer response.Body.Close()

	fmt.Println(response.StatusCode)

	if response.StatusCode != 200 {
		errorMsg := fmt.Sprintf("error get payment methods, status code: %d", response.StatusCode)
		return resp, errors.New(errorMsg)
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
