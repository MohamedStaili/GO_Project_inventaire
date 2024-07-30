package main

import (
	"MohamedStaili/GO_Project_inventaire/pkg/models"
	"MohamedStaili/GO_Project_inventaire/pkg/routes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	//supprimer les sessions expirés
	go func() {
		if err := models.DeleteExpiredSessions(); err != nil {
			fmt.Printf("error not deleted session:%v", err)
		} else {
			fmt.Println("session deleted succefly")
		}
		time.Sleep(1 * time.Hour)
	}()
	router := mux.NewRouter().StrictSlash(true)
	routes.RegisterInventaireRoutes(router)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"}) // Remplacez "*" par le domaine de votre frontend en production
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)))
}
