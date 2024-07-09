package routes

import (
	"github.com/gorilla/mux"
)

var RegisterInventaireRoutes = func(router *mux.Router) {
	// Routes pour l'administration des utilisateurs
	router.HandleFunc("/admin/user", controllers.AddUser).Methods("POST")           // Ajouter un utilisateur (admin)
	router.HandleFunc("/admin/user/{id}", controllers.DeleteUser).Methods("DELETE") // Supprimer un utilisateur par ID (admin)

	// Autres routes pour les pages spécifiques
	router.HandleFunc("/login", controllers.LoginPage).Methods("GET")                       // Page de connexion (admin + utilisateur)
	router.HandleFunc("/profile", controllers.ProfilePage).Methods("GET")                   // Page de profil
	router.HandleFunc("/inventaire/add", controllers.AjouterInventaire).Methods("GET")      // Page d'ajout d'inventaire
	router.HandleFunc("/inventaire/delete", controllers.SupprimerInventaire).Methods("GET") // Page de suppression d'inventaire
	router.HandleFunc("/inventaire/modify", controllers.ModifierInventaire).Methods("GET")  // Page de modification d'inventaire
	router.HandleFunc("/search", controllers.SearchPage).Methods("GET")                     // Page de recherche
	router.HandleFunc("/admin/adduser", controllers.AjouterUtilisateur).Methods("GET")      // Page d'ajout d'utilisateur (admin)
	router.HandleFunc("/admin/deleteuser", controllers.SupprimerUtilisateur).Methods("GET") // Page de suppression d'utilisateur (admin)
	router.HandleFunc("/home", controllers.HomePage).Methods("GET")                         // Page d'accueil avec des statistiques sur les inventaires

}
