package vehicle

type Vehicle struct {
	ID    int    `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

func NewCar(brand, model string, year, lastID int) Vehicle {
	return Vehicle{
		ID:    lastID + 1,
		Brand: brand,
		Model: model,
		Year:  year,
	}
}
