package main

import (
	"encoding/json"
	"fmt"
	"go-highschool-api/src/handler"
	"net/http"
	"os"
)

func main() {
	port := ":8080"

	fmt.Printf("Server starting on port %s\n", port)

	http.HandleFunc("/cars", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handler.ListVehicles(w, r)

		case "POST":
			handler.CreateVehicle(w, r)

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

		handler.RetrieveVehicle(w, r, r.PathValue("id"))
	})

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
