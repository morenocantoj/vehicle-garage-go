package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Vehiculo struct {
	identificador int
	Marca         string
	Modelo        string
	Anyo          int
}

func main() {
	puerto := ":8080"

	fmt.Printf("Servidor iniciando en el puerto %s\n", puerto)

	coches := []Vehiculo{
		{identificador: 1, Marca: "Toyota", Modelo: "Corolla", Anyo: 2020},
		{identificador: 2, Marca: "Honda", Modelo: "Civic", Anyo: 2019},
		{identificador: 3, Marca: "Ford", Modelo: "Mustang", Anyo: 2021},
	}

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
