package handler

import (
	"encoding/json"
	"fmt"
	"go-highschool-api/src/vehicle"
	"net/http"

	"github.com/google/uuid"
)

type VehicleHandler struct {
	repository vehicle.Repository
}

func NewVehicleHandler(repository vehicle.Repository) *VehicleHandler {
	return &VehicleHandler{
		repository: repository,
	}
}

type CarRequest struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

func (h *VehicleHandler) ListVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	cars, err := h.repository.FindAll()
	if err != nil {
		fmt.Printf("Error retrieving cars %w\n", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Internal server error",
		})
		return
	}

	json.NewEncoder(w).Encode(cars)
}

func (h *VehicleHandler) CreateVehicle(w http.ResponseWriter, r *http.Request) {
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

	newCar := vehicle.NewCar(carRequest.Brand, carRequest.Model, carRequest.Year)
	err = h.repository.Save(newCar)
	if err != nil {
		fmt.Printf("Error saving car %w\n", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Internal server error",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCar)
}

func (h *VehicleHandler) RetrieveVehicle(w http.ResponseWriter, r *http.Request, rawIdentifier string) {
	id, err := uuid.Parse(rawIdentifier)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid ID format",
		})
		return
	}

	car, err := h.repository.FindByID(id)
	if err != nil {
		fmt.Printf("Error retrieving car %w\n", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Internal server error",
		})
		return
	}

	if car == (vehicle.Vehicle{}) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Car not found",
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(car)
}
