package handler

import (
	"encoding/json"
	"go-highschool-api/src/vehicle"
	"net/http"
	"strconv"
)

type CarRequest struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

var cars = []vehicle.Vehicle{}

func ListVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cars)
}

func CreateVehicle(w http.ResponseWriter, r *http.Request) {
	var carRequest CarRequest
	err := json.NewDecoder(r.Body).Decode(&carRequest)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid JSON format",
		})
		return
	}

	newCar := vehicle.NewCar(carRequest.Brand, carRequest.Model, carRequest.Year, len(cars))
	cars = append(cars, newCar)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCar)
}

func RetrieveVehicle(w http.ResponseWriter, r *http.Request, rawIdentifier string) {
	id, err := strconv.Atoi(rawIdentifier)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid ID format",
		})
		return
	}

	for _, car := range cars {
		if car.ID == id {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(car)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "Car not found",
	})
}
