package routes

import (
	"MohamedStaili/GO_Project_inventaire/pkg/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

var RegisterInventaireRoutes = func(router *mux.Router) {

	// Servir des fichiers statiques
	staticDir := "/static/"
	router.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	// Routes pour l'administration des utilisateurs
	router.HandleFunc("/admin/user/adduser", controllers.AjouterUser).Methods("POST")            // Page d'ajout d'utilisateur (admin)
	router.HandleFunc("/admin/user/deleteuser{id}", controllers.SupprimerUser).Methods("DELETE") // Page de suppression d'utilisateur (admin)

	// Autres routes pour les pages spécifiques
	router.HandleFunc("/login", controllers.LoginPage).Methods("POST") // Page de connexion (admin + utilisateur)

	router.HandleFunc("/user/logout", controllers.LogOut).Methods("POST") // Log out

	router.HandleFunc("/inventaire/add", controllers.AjouterInventaire).Methods("POST")            // Page d'ajout d'inventaire
	router.HandleFunc("/inventaire/delete{id}", controllers.SupprimerInventaire).Methods("DELETE") // Page de suppression d'inventaire
	router.HandleFunc("/inventaire/modify", controllers.ModifierInventaire).Methods("PUT")         // Page de modification d'inventaire
	router.HandleFunc("/search{id}", controllers.SearchPage).Methods("GET")                        // Page de recherche

	router.HandleFunc("/home", controllers.HomePage).Methods("GET") // Page d'accueil avec des statistiques sur les inventaires
}
