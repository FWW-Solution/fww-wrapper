package adapter

import (
	"fww-wrapper/internal/data/dto_booking"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/goccy/go-json"
)

// Booking implements Adapter.
func (a *adapter) Booking(body *dto_booking.Request) (resp dto_booking.AsyncBookResponse, err error) {
	json, err := json.Marshal(body)
	if err != nil {
		return resp, err
	}

	ID := watermill.NewUUID()

	err = a.publisher.Publish("request_booking", message.NewMessage(
		ID,
		json,
	))
	if err != nil {
		return resp, err
	}

	resp = dto_booking.AsyncBookResponse{
		BookingIDCode: ID,
	}

	return resp, nil

}
