package dto_airport

type ResponseAirport struct {
	City      string `mapstructure:"city"`
	CreatedAt string `mapstructure:"created_at"`
	Iata      string `mapstructure:"iata"`
	Icao      string `mapstructure:"icao"`
	ID        int64  `mapstructure:"id"`
	Name      string `mapstructure:"name"`
	Province  string `mapstructure:"province"`
	UpdatedAt string `mapstructure:"updated_at"`
}
