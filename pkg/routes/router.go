package routes

import (
	"MohamedStaili/GO_Project_inventaire/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterInventaireRoutes = func(router *mux.Router) {
	// Routes pour l'administration des utilisateurs
	//router.HandleFunc("/admin/user/adduser", controllers.AjouterUtilisateur).Methods("Post")            // Page d'ajout d'utilisateur (admin)
	//router.HandleFunc("/admin/user/deleteuser{id}", controllers.SupprimerUtilisateur).Methods("DELETE") // Page de suppression d'utilisateur (admin)
	// Autres routes pour les pages spécifiques
	//router.HandleFunc("/login", controllers.LoginPage).Methods("GET")                              // Page de connexion (admin + utilisateur)
	//router.HandleFunc("/profile", controllers.ProfilePage).Methods("GET")                          // Page de profil
	router.HandleFunc("/inventaire/add", controllers.AjouterInventaire).Methods("POST")            // Page d'ajout d'inventaire
	router.HandleFunc("/inventaire/delete{id}", controllers.SupprimerInventaire).Methods("DELETE") // Page de suppression d'inventaire
	router.HandleFunc("/inventaire/modify", controllers.ModifierInventaire).Methods("PUT")         // Page de modification d'inventaire
	router.HandleFunc("/search{id}", controllers.SearchPage).Methods("GET")                        // Page de recherche
	//router.HandleFunc("/home", controllers.HomePage).Methods("GET")                                // Page d'accueil avec des statistiques sur les inventaires

}
