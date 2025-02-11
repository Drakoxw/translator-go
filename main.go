package main

import (
	"log"
	"net/http"

	"translate/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/translate", handlers.TranslateHandler).Methods("POST", "PUT")

	log.Println("Servidor iniciado en el puerto 8080")
	http.ListenAndServe(":8080", r)
}
