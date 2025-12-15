package vehicle

import "github.com/google/uuid"

type Vehicle struct {
	ID    uuid.UUID `json:"id"`
	Brand string    `json:"brand"`
	Model string    `json:"model"`
	Year  int       `json:"year"`
}

func NewCar(brand, model string, year int) Vehicle {
	return Vehicle{
		ID:    uuid.New(),
		Brand: brand,
		Model: model,
		Year:  year,
	}
}
