package dto_flight

type ResponseFlightDetail struct {
	ArrivalAirportName  string  `mapstructure:"arrival_airport_name"`
	ArrivalTime         string  `mapstructure:"arrival_time"`
	CodeFlight          string  `mapstructure:"code_flight"`
	DepartureTime       string  `mapstructure:"departure_time"`
	DepatureAirportName string  `mapstructure:"depature_airport_name"`
	FlightPrice         float64 `mapstructure:"flight_price"`
	ReminingSeat        int     `mapstructure:"remining_seat"`
	Status              string  `mapstructure:"status"`
}

type ResponseFlight struct {
	ArrivalAirportName  string  `mapstructure:"arrival_airport_name"`
	ArrivalTime         string  `mapstructure:"arrival_time"`
	CodeFlight          string  `mapstructure:"code_flight"`
	DepartureTime       string  `mapstructure:"departure_time"`
	DepatureAirportName string  `mapstructure:"depature_airport_name"`
	FlightPrice         float64 `mapstructure:"flight_price"`
	ReminingSeat        int     `mapstructure:"remining_seat"`
	Status              string  `mapstructure:"status"`
}
