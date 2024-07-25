package routes

import (
	"MohamedStaili/GO_Project_inventaire/middlewares"
	"MohamedStaili/GO_Project_inventaire/pkg/controllers"

	"net/http"

	"github.com/gorilla/mux"
)

var RegisterInventaireRoutes = func(router *mux.Router) {

	// Servir des fichiers statiques
	/*staticDir := "/static/"
	router.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))*/
	// Middleware pour vérifier l'authentification
	router.PathPrefix("/static/").Handler(middlewares.AuthRequired(http.StripPrefix("/static/", http.FileServer(http.Dir("static/")))))

	// Autres routes pour les pages spécifiques
	router.HandleFunc("/login", controllers.LoginPage).Methods("POST") // Page de connexion (admin + utilisateur)

	router.HandleFunc("/user/logout", controllers.LogOut).Methods("POST") // Log out

	router.HandleFunc("/inventaire/add", controllers.AjouterInventaire).Methods("POST")            // Page d'ajout d'inventaire
	router.HandleFunc("/inventaire/delete{id}", controllers.SupprimerInventaire).Methods("DELETE") // Page de suppression d'inventaire
	router.HandleFunc("/inventaire/modify", controllers.ModifierInventaire).Methods("PUT")         // Page de modification d'inventaire
	router.HandleFunc("/search{id}", controllers.SearchPage).Methods("GET")                        // Page de recherche

	router.HandleFunc("/GEtInventory", controllers.GetAllInventaire).Methods("GET") // Page d'accueil avec des statistiques sur les inventaires
	// Routes accessibles uniquement aux administrateurs
	admin := router.PathPrefix("/admin").Subrouter()
	admin.Use(middlewares.AdminOnlyMiddleware)
	admin.HandleFunc("/adduser", controllers.AjouterUser).Methods("POST")
	//admin.HandleFunc("/userlist", controllers.UserList).Methods("GET")
	admin.HandleFunc("/deleteuser", controllers.SupprimerUser).Methods("DELETE")
}
