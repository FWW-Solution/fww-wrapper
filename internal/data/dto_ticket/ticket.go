package dto_ticket

type Request struct {
	CodeBooking string `json:"code_booking"`
}

type Response struct {
	BordingTime string `mapstructure:"bording_time"`
	CodeTicket  string `mapstructure:"code_ticket"`
}
