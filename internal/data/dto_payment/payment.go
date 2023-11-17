package dto_payment

type Request struct {
	BookingID     int64  `json:"booking_id"`
	PaymentMethod string `json:"payment_method"`
}

type AsyncPaymentResponse struct {
	PaymentIDCode string `json:"payment_id_code"`
}

type StatusResponse struct {
	Status string `mapstructure:"status"`
}

type MethodResponse struct {
	ID       int64  `mapstructure:"id"`
	IsActive bool   `mapstructure:"is_active"`
	Name     string `mapstructure:"name"`
}
