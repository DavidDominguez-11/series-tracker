package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rs/cors" // Importa el paquete CORS

	"series-tracker/handlers"
	"series-tracker/db"
)

func main() {
	// Iniciar conexión a la base de datos
	db.InitDB()
	defer db.CloseDB()

	// Inicializar el router
	r := mux.NewRouter()

	// Configuración de CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// Rutas
	r.HandleFunc("/api/series", handlers.GetSeries).Methods("GET")
	r.HandleFunc("/api/series/{id}", handlers.GetSeriesByID).Methods("GET")
	r.HandleFunc("/api/series", handlers.CreateSeries).Methods("POST")
	r.HandleFunc("/api/series/{id}", handlers.UpdateSeries).Methods("PUT")
	r.HandleFunc("/api/series/{id}", handlers.DeleteSeries).Methods("DELETE")

	// Envolver el router con CORS
	handler := c.Handler(r)

	// Iniciar el servidor
	//log.Println("Servidor corriendo en puerto 8080...")
	//http.ListenAndServe(":8080", r)
	log.Println("Servidor corriendo en el puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
