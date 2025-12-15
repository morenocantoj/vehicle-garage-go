package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Vehicle struct {
	id    int
	Brand string
	Model string
	Year  int
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

	http.HandleFunc("/view-car/{id}", func(w http.ResponseWriter, r *http.Request) {
		carIDParam := r.PathValue("id")
		if carIDParam == "" {
			w.Write([]byte("Missing car ID"))
			return
		}

		carID, err := strconv.Atoi(carIDParam)
		if err != nil {
			w.Write([]byte(fmt.Errorf("car ID is not a valid number: %w", err).Error()))
			return
		}

		for _, car := range cars {
			if car.id == carID {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(car)
				return
			}
		}

		w.Write([]byte("Car not found"))
	})

	http.HandleFunc("/create-car", func(w http.ResponseWriter, r *http.Request) {
		var carRequest CarRequest
		err := json.NewDecoder(r.Body).Decode(&carRequest)
		if err != nil {
			w.Write([]byte(fmt.Errorf("failed to read request: %w", err).Error()))
			return
		}

		newCar := createCar(carRequest.Brand, carRequest.Model, carRequest.Year, len(cars))
		cars = append(cars, newCar)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newCar)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]any{
			"message": "List of cars",
			"cars":    cars,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
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
		id:    lastID + 1,
		Brand: brand,
		Model: model,
		Year:  year,
	}
}
