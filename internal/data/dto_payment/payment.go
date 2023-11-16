package dto_payment

type Request struct {
	BookingID     string `json:"booking_id"`
	PaymentMethod string `json:"payment_method"`
}

type AsyncPaymentResponse struct {
	PaymentIDCode string `json:"payment_id_code"`
}

type StatusResponse struct {
	Status string `mapstructure:"status"`
}