package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Vehicle struct {
	ID    int    `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

type CarRequest struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

func main() {
	port := ":8080"

	fmt.Printf("Server starting on port %s\n", port)

	cars := []Vehicle{}

	http.HandleFunc("/cars", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(cars)

		case "POST":
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

			newCar := createCar(carRequest.Brand, carRequest.Model, carRequest.Year, len(cars))
			cars = append(cars, newCar)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newCar)

		default:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Method not allowed",
			})
		}
	})

	http.HandleFunc("/cars/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Method not allowed",
			})
			return
		}

		carIDParam := r.PathValue("id")
		if carIDParam == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Missing car ID",
			})
			return
		}

		carID, err := strconv.Atoi(carIDParam)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Invalid car ID format",
			})
			return
		}

		for _, car := range cars {
			if car.ID == carID {
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
	})

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func createCar(brand, model string, year, lastID int) Vehicle {
	return Vehicle{
		ID:    lastID + 1,
		Brand: brand,
		Model: model,
		Year:  year,
	}
}
