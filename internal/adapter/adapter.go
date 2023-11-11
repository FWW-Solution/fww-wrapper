package adapter

import (
	"fww-wrapper/internal/config"
	"fww-wrapper/internal/data/dto_airport"
	"fww-wrapper/internal/data/dto_booking"
	"fww-wrapper/internal/data/dto_flight"
	"fww-wrapper/internal/data/dto_passanger"

	"github.com/ThreeDotsLabs/watermill/message"
	circuit "github.com/rubyist/circuitbreaker"
)

type adapter struct {
	client    *circuit.HTTPClient
	cfg       *config.HttpClientConfig
	publisher message.Publisher
}

type Adapter interface {
	GetPassanger(id int) (resp dto_passanger.ResponseDetail, err error)
	RegisterPassanger(body *dto_passanger.RequestRegister) (resp dto_passanger.ResponseRegistered, err error)
	UpdatePassanger(body *dto_passanger.RequestUpdate) (resp dto_passanger.ResponseUpdate, err error)
	GetAirport(city, province, iata string) (resp []dto_airport.ResponseAirport, err error)
	GetFlights(departureTime, arrivalTime string, limit int, offset int) (resp []dto_flight.ResponseFlight, err error)
	GetDetailFlightByID(id int64) (resp dto_flight.ResponseFlightDetail, err error)
	// Booking
	Booking(body *dto_booking.Request) (resp dto_booking.AsyncBookResponse, err error)
}

func New(client *circuit.HTTPClient, cfg *config.HttpClientConfig, publisher message.Publisher) Adapter {
	return &adapter{
		client:    client,
		cfg:       cfg,
		publisher: publisher,
	}
}
