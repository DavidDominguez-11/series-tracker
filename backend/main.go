package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"series-tracker/handlers"
	"series-tracker/db"
)

func main() {
	// Iniciar conexión a la base de datos
	db.InitDB()
	defer db.CloseDB()

	// Inicializar el router
	r := mux.NewRouter()

	// Rutas
	r.HandleFunc("/api/series", handlers.GetSeries).Methods("GET")
	r.HandleFunc("/api/series/{id:[0-9]+}", handlers.GetSeriesByID).Methods("GET")
	r.HandleFunc("/api/series", handlers.CreateSeries).Methods("POST")
	r.HandleFunc("/api/series/{id:[0-9]+}", handlers.UpdateSeries).Methods("PUT")
	r.HandleFunc("/api/series/{id:[0-9]+}", handlers.DeleteSeries).Methods("DELETE")

	// Iniciar el servidor
	log.Println("Servidor corriendo en puerto 8080...")
	http.ListenAndServe(":8080", r)
}
