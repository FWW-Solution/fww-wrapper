package dto_booking

type Request struct {
	BookDetails []BookDetail `json:"book_details"`
	FlightID    int64        `json:"flight_id"`
	UserID      int64        `json:"user_id"`
}

type BookDetail struct {
	Baggage     int     `json:"baggage"`
	Class       string  `json:"class"`
	PassangerID int64   `json:"passanger_id"`
	SeatNumber  *string `json:"seat_number"`
}

type AsyncBookResponse struct {
	BookingIDCode string `json:"booking_id_code"`
}

type BookResponse struct {
	// Airport Name
	ArrivalAirport string `mapstructure:"arrival_airport"`
	ArrivalTime    string `mapstructure:"arrival_time"`
	BookExpiredAt  string `mapstructure:"book_expired_at"`
	CodeBooking    string `mapstructure:"code_booking"`
	CodeFlight     string `mapstructure:"code_flight"`
	// Airport Name
	DepartureAirport string               `mapstructure:"departure_airport"`
	DepartureTime    string               `mapstructure:"departure_time"`
	Details          []BookResponseDetail `mapstructure:"details"`
	ID               int64                `mapstructure:"id"`
	PaymentExpiredAt string               `mapstructure:"payment_expired_at"`
	TotalPrice       float64              `mapstructure:"total_price"`
}

type BookResponseDetail struct {
	Bagage        int     `mapstructure:"bagage"`
	Class         string  `mapstructure:"class"`
	PassangerName string  `mapstructure:"passanger_name"`
	Price         float64 `mapstructure:"price"`
	SeatNumber    string  `mapstructure:"seat_number"`
}
