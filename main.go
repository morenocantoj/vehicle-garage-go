package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Vehiculo struct {
	identificador int
	Marca         string
	Modelo        string
	Anyo          int
}

type cocheRequest struct {
	Marca  string `json:"marca"`
	Modelo string `json:"modelo"`
	Anyo   int    `json:"anyo"`
}

func main() {
	puerto := ":8080"

	fmt.Printf("Servidor iniciando en el puerto %s\n", puerto)

	coches := []Vehiculo{}

	http.HandleFunc("/ver-coche/{id}", func(w http.ResponseWriter, r *http.Request) {
		identificadorCocheParam := r.PathValue("id")
		if identificadorCocheParam == "" {
			w.Write([]byte("Falta el identificador del coche"))
			return
		}

		identificadorCoche, err := strconv.Atoi(identificadorCocheParam)
		if err != nil {
			w.Write([]byte(fmt.Errorf("identificador del coche no es un número válido: %w", err).Error()))
			return
		}

		for _, coche := range coches {
			if coche.identificador == identificadorCoche {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(coche)
				return
			}
		}

		w.Write([]byte("Coche no encontrado"))
	})

	http.HandleFunc("/crear-coche", func(w http.ResponseWriter, r *http.Request) {
		var cocheBytes cocheRequest
		err := json.NewDecoder(r.Body).Decode(&cocheBytes)
		if err != nil {
			w.Write([]byte(fmt.Errorf("fallo al leer la request %w", err).Error()))
			return
		}

		nuevoCoche := crearCoche(cocheBytes.Marca, cocheBytes.Modelo, cocheBytes.Anyo, len(coches))
		coches = append(coches, nuevoCoche)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(nuevoCoche)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		respuesta := map[string]any{
			"mensaje": "Lista de coches",
			"coches":  coches,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(respuesta)
	})

	err := http.ListenAndServe(puerto, nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func crearCoche(marca, modelo string, anyo, ultimoIdentificador int) Vehiculo {
	return Vehiculo{
		identificador: ultimoIdentificador + 1,
		Marca:         marca,
		Modelo:        modelo,
		Anyo:          anyo,
	}
}
