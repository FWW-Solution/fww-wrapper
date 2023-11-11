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
	ArrivalAirport string `json:"arrival_airport"`
	ArrivalTime    string `json:"arrival_time"`
	BookExpiredAt  string `json:"book_expired_at"`
	CodeBooking    string `json:"code_booking"`
	CodeFlight     string `json:"code_flight"`
	// Airport Name
	DepartureAirport string               `json:"departure_airport"`
	DepartureTime    string               `json:"departure_time"`
	Details          []BookResponseDetail `json:"details"`
	ID               int64                `json:"id"`
	PaymentExpiredAt string               `json:"payment_expired_at"`
	TotalPrice       float64              `json:"total_price"`
}

type BookResponseDetail struct {
	Bagage        int     `json:"bagage"`
	Class         string  `json:"class"`
	PassangerName string  `json:"passanger_name"`
	Price         float64 `json:"price"`
	SeatNumber    string  `json:"seat_number"`
}
