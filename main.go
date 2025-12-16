package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	puerto := ":8080"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Ye que pasa mundo"))
	})

	fmt.Printf("Servidor iniciando en el puerto %s\n", puerto)

	err := http.ListenAndServe(puerto, nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
