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

func (v Vehicle) ContaminationBadge() string {
	switch {
	case v.Year >= 2020:
		return "A"
	case v.Year >= 2015:
		return "B"
	case v.Year >= 2010:
		return "C"
	case v.Year >= 2005:
		return "D"
	default:
		return "E"
	}
}
