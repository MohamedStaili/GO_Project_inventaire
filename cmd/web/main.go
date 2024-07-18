package main

import (
	"MohamedStaili/GO_Project_inventaire/pkg/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterInventaireRoutes(r)
	http.Handle("/", r)
	/*headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"}) // Remplacez "*" par le domaine de votre frontend en production
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(r)))*/
}
