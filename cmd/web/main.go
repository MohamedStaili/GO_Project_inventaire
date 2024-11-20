package main

import (
	"MohamedStaili/GO_Project_inventaire/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	routes.RegisterInventaireRoutes(router)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:3000"}) // Remplacez par le domaine de votre frontend

	// Middleware CORS avec support pour les credentials
	corsHandler := handlers.CORS(
		headers,
		methods,
		origins,
		handlers.AllowCredentials(), // Permet les credentials (cookies)
	)

	log.Fatal(http.ListenAndServe(":8080", corsHandler(router)))
}
