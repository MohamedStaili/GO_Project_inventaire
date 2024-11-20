package routes

import (
	"MohamedStaili/GO_Project_inventaire/middlewares"
	"MohamedStaili/GO_Project_inventaire/pkg/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

var RegisterInventaireRoutes = func(router *mux.Router) {

	// Public routes (no authentication required)
	router.HandleFunc("/login", controllers.LoginPage).Methods("POST")    // Login page
	router.HandleFunc("/user/logout", controllers.LogOut).Methods("POST") // Logout

	// Authenticated user routes (authentication required)
	user := router.PathPrefix("/user").Subrouter()
	user.Use(middlewares.AuthRequired)                                      // Apply AuthRequired middleware
	user.HandleFunc("/getuserinfo", controllers.GetUserInfo).Methods("GET") // Get user info
	// Assuming you have a specific middleware function called SpecificMiddleware
	user.Handle("/profile", middlewares.AuthProfile((http.HandlerFunc(controllers.ModifierProfile)))).Methods("PUT")

	// Inventory management routes (authentication required)
	inventory := router.PathPrefix("/inventaire").Subrouter()
	inventory.Use(middlewares.AuthRequired)                                                // Apply AuthRequired middleware
	inventory.HandleFunc("/add", controllers.AjouterInventaire).Methods("POST")            // Add inventory
	inventory.HandleFunc("/delete{id}", controllers.SupprimerInventaire).Methods("DELETE") // Delete inventory
	inventory.HandleFunc("/modify", controllers.ModifierInventaire).Methods("PUT")         // Modify inventory
	inventory.HandleFunc("/search{id}", controllers.SearchPage).Methods("GET")             // Search inventory
	inventory.HandleFunc("/list", controllers.GetAllInventaire).Methods("GET")             // List all inventory items
	materiel := router.PathPrefix("/materiel").Subrouter()
	//materiel api
	materiel.Use(middlewares.AuthRequired)                                              // Apply AuthRequired middleware
	materiel.HandleFunc("/add", controllers.AjouterMateriel).Methods("POST")            // Add materiel
	materiel.HandleFunc("/delete{id}", controllers.SupprimerMateriel).Methods("DELETE") // Delete materiel
	materiel.HandleFunc("/modify", controllers.ModifierMateriel).Methods("PUT")         // Modify materiel
	//materiel.HandleFunc("/search{id}", controllers.SearchPage).Methods("GET")             // Search materiel
	materiel.HandleFunc("/list", controllers.GetAllMateriel).Methods("GET") // List all inventory items
	//employee api
	employee := router.PathPrefix("/employee").Subrouter()
	employee.Use(middlewares.AuthRequired)                                             // Apply AuthRequired middleware
	employee.HandleFunc("/add", controllers.AjouterEmployee).Methods("POST")           // Add employee
	employee.HandleFunc("/delete{id}", controllers.SupprimerEmploye).Methods("DELETE") // Delete employee
	employee.HandleFunc("/modify", controllers.ModifierEmploye).Methods("PUT")         // Modify employee
	//employee.HandleFunc("/search{id}", controllers.SearchPage).Methods("GET")             // Search employee
	employee.HandleFunc("/list", controllers.GetAllEmployee).Methods("GET") // List all inventory items
	//achat api
	achat := router.PathPrefix("/purchase").Subrouter()
	achat.Use(middlewares.AuthRequired)                                           // Apply AuthRequired middleware
	achat.HandleFunc("/add", controllers.AjouterAchat).Methods("POST")            // Add achat
	achat.HandleFunc("/delete{id}", controllers.SupprimerAchat).Methods("DELETE") // Delete achat
	achat.HandleFunc("/modify", controllers.ModifierAchat).Methods("PUT")         // Modify achat
	//achat.HandleFunc("/search{id}", controllers.SearchPage).Methods("GET")             // Search achat
	achat.HandleFunc("/list", controllers.GetAllPurchase).Methods("GET") // List all inventory items

	// Admin routes (admin-only access)
	admin := router.PathPrefix("/admin").Subrouter()
	admin.Use(middlewares.AdminOnlyMiddleware)                                       // Apply AdminOnlyMiddleware
	admin.HandleFunc("/adduser", controllers.AjouterUser).Methods("POST")            // Add user (admin only)
	admin.HandleFunc("/deleteuser{id}", controllers.SupprimerUser).Methods("DELETE") // Delete user (admin only)
	admin.HandleFunc("/GetListUser", controllers.GetListUser).Methods("POST")        // Delete user (admin only)

	// Optionally: Static files route (with or without AuthRequired)
	// router.PathPrefix("/static/").Handler(middlewares.AuthRequired(http.StripPrefix("/static/", http.FileServer(http.Dir("static/")))))
}
