package dto_passanger

type RequestRegister struct {
	FullName    string `json:"full_name" validate:"required"`
	Gender      string `json:"gender" validate:"required"`
	IDNumber    string `json:"id_number" validate:"required,min=16,max=16"`
	IDType      string `json:"id_type" validate:"required"`
	DateOfBirth string `json:"date_of_birth" validate:"required,datetime=2006-01-02"`
}

type RequestUpdate struct {
	DateOfBirth string `json:"date_of_birth" validate:"required,datetime=2006-01-02"`
	FullName    string `json:"full_name" validate:"required"`
	Gender      string `json:"gender" validate:"required"`
	ID          int64  `json:"id" validate:"required,numeric"`
	IDNumber    string `json:"id_number" validate:"required,min=16,max=16"`
	IDType      string `json:"id_type" validate:"required"`
}

type ResponseRegistered struct {
	ID int64 `json:"id"`
}

type ResponseDetail struct {
	CovidVaccineStatus string `mapstructure:"covid_vaccine_status"`
	CreatedAt          string `mapstructure:"created_at"`
	DateOfBirth        string `mapstructure:"date_of_birth"`
	FullName           string `mapstructure:"full_name"`
	Gender             string `mapstructure:"gender"`
	ID                 int64  `mapstructure:"id"`
	IDNumber           string `mapstructure:"id_number"`
	IDType             string `mapstructure:"id_type"`
	IsIDVerified       bool   `mapstructure:"is_id_verified"`
	UpdatedAt          string `mapstructure:"updated_at"`
}

type ResponseUpdate struct {
	ID int64 `mapstructure:"id"`
}
