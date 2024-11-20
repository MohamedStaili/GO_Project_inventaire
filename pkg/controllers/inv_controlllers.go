package controllers

import (
	"MohamedStaili/GO_Project_inventaire/pkg/config"
	"MohamedStaili/GO_Project_inventaire/pkg/models"
	utils "MohamedStaili/GO_Project_inventaire/pkg/utiles"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func SearchPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error: in parseint function")
	}
	bookDetails, _ := models.SearchPage(Id)
	//config server part
	res, _ := json.Marshal(bookDetails) //transform data to json form
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) //code=200
	w.Write(res)
}

func AjouterInventaire(w http.ResponseWriter, r *http.Request) {
	var NewInv models.Inventaire
	utils.ParseBody(r, &NewInv)

	// Ajouter l'inventaire en vérifiant les clés étrangères
	b, err := NewInv.AjouterInventaire()
	if err != nil {
		// Si une erreur est survenue, retourner une erreur JSON avec un statut HTTP approprié
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Si tout va bien, retourner le nouvel inventaire en JSON
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func AjouterEmployee(w http.ResponseWriter, r *http.Request) {
	var NewEmpl models.Employe
	utils.ParseBody(r, &NewEmpl)

	// Ajouter l'inventaire en vérifiant les clés étrangères
	b, err := NewEmpl.AjouterEmployee()
	if err != nil {
		// Si une erreur est survenue, retourner une erreur JSON avec un statut HTTP approprié
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Si tout va bien, retourner le nouvel inventaire en JSON
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func AjouterMateriel(w http.ResponseWriter, r *http.Request) {
	var NewMat models.Materiel
	utils.ParseBody(r, &NewMat)

	// Ajouter l'inventaire en vérifiant les clés étrangères
	b, err := NewMat.AjouterMateriel()
	if err != nil {
		// Si une erreur est survenue, retourner une erreur JSON avec un statut HTTP approprié
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Si tout va bien, retourner le nouvel inventaire en JSON
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func AjouterAchat(w http.ResponseWriter, r *http.Request) {
	var NewAchat models.Achat
	utils.ParseBody(r, &NewAchat)

	// Ajouter l'inventaire en vérifiant les clés étrangères
	b, err := NewAchat.AjouterAchat()
	if err != nil {
		// Si une erreur est survenue, retourner une erreur JSON avec un statut HTTP approprié
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Si tout va bien, retourner le nouvel inventaire en JSON
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func SupprimerInventaire(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	invId := vars["id"]
	Id, err := strconv.ParseInt(invId, 0, 0)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = models.SupprimerInventaire(Id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			http.Error(w, "Record not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to delete record", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Record deleted successfully"}`))
}
func SupprimerMateriel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	invId := vars["id"]
	Id, err := strconv.ParseInt(invId, 0, 0)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = models.SupprimerMateriel(Id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			http.Error(w, "Record not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to delete record", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Record deleted successfully"}`))
}
func SupprimerEmploye(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	invId := vars["id"]
	Id, err := strconv.ParseInt(invId, 0, 0)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = models.SupprimerEmploye(Id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			http.Error(w, "Record not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to delete record", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Record deleted successfully"}`))
}
func SupprimerAchat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	invId := vars["id"]
	Id, err := strconv.ParseInt(invId, 0, 0)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = models.SupprimerAchat(Id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			http.Error(w, "Record not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to delete record", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Record deleted successfully"}`))
}

// Handler HTTP pour modifier un inventaire
func ModifierInventaire(w http.ResponseWriter, r *http.Request) {
	// Créer une instance de modèle Inventaire pour stocker les données de la requête
	modInv := models.Inventaire{}

	// Parser le corps de la requête JSON dans modInv
	utils.ParseBody(r, &modInv)

	// Appeler la méthode ModifierInventaire() pour mettre à jour l'inventaire
	modInv.ModifierInventaire()

	// Répondre avec le résultat JSON
	res, err := json.Marshal(modInv)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
func ModifierMateriel(w http.ResponseWriter, r *http.Request) {
	// Créer une instance de modèle ModifierMateriel pour stocker les données de la requête
	modInv := models.Materiel{}

	// Parser le corps de la requête JSON dans modInv
	utils.ParseBody(r, &modInv)

	// Appeler la méthode ModifierMateriel() pour mettre à jour l'MModifierMateriel
	modInv.ModifierMateriel()

	// Répondre avec le résultat JSON
	res, err := json.Marshal(modInv)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
func ModifierEmploye(w http.ResponseWriter, r *http.Request) {
	// Créer une instance de modèle Inventaire pour stocker les données de la requête
	modInv := models.Employe{}

	// Parser le corps de la requête JSON dans modInv
	utils.ParseBody(r, &modInv)

	// Appeler la méthode ModifierInventaire() pour mettre à jour l'inventaire
	modInv.ModifierEmploye()

	// Répondre avec le résultat JSON
	res, err := json.Marshal(modInv)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
func ModifierAchat(w http.ResponseWriter, r *http.Request) {
	// Créer une instance de modèle Inventaire pour stocker les données de la requête
	modInv := models.Achat{}

	// Parser le corps de la requête JSON dans modInv
	utils.ParseBody(r, &modInv)

	// Appeler la méthode ModifierInventaire() pour mettre à jour l'inventaire
	modInv.ModifierAchat()

	// Répondre avec le résultat JSON
	res, err := json.Marshal(modInv)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// Create a User
func AjouterUser(w http.ResponseWriter, r *http.Request) {
	var NewUser models.User
	utils.ParseBody(r, &NewUser)
	fmt.Printf("Parsed User: %+v\n", NewUser)
	//verifier l'existence de{username,password,email}
	if NewUser.Username == "" || NewUser.Email == "" || NewUser.Password == "" || NewUser.Role == "" {
		http.Error(w, "plaise provide all fields", http.StatusBadRequest)
		return
	}
	user, err := NewUser.AjouterUser()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func ModifierProfile(w http.ResponseWriter, r *http.Request) {
	modUser := models.User{}

	// Parse the request body
	if err := utils.ParseBody(r, &modUser); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Extract user ID from the context
	userIDFromToken, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "Invalid user ID in context", http.StatusUnauthorized)
		return
	}
	fmt.Println(userIDFromToken)
	// Convert userIDFromToken from string to uint
	userID, err := strconv.ParseUint(userIDFromToken, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}
	fmt.Println(userID)

	// Check if the user is trying to modify their own profile
	modUser.ID = uint(userID)
	fmt.Println(modUser.ID)
	// Proceed to modify the profile
	updatedUser, err := modUser.ModifierProfile()
	if err != nil {
		http.Error(w, "Error updating profile", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(updatedUser)
	if err != nil {
		http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func SupprimerUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]
	UserId, err := strconv.ParseInt(ID, 0, 0)
	if err != nil {
		fmt.Printf("Error Format:%v", err)
		return
	}
	err = models.SupprimerUser(UserId)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			http.Error(w, "Record not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to delete record", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Record deleted successfully"}`))
}

var SecretKey = "secret"

func LoginPage(w http.ResponseWriter, r *http.Request) {
	var LogInfo map[string]string
	err := utils.ParseBody(r, &LogInfo)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	email := LogInfo["email"]
	password := LogInfo["password"]
	var user models.User

	err, user = models.LoginPage(email, password)
	if err != nil {
		fmt.Printf("error :%v", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// JWT authentication
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	// Create the token
	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set JWT as a cookie
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	res, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func LogOut(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    "",                         // Clear the value
		Expires:  time.Now().Add(-time.Hour), // Set expiration to a past time
		Path:     "/",                        // Ensure the path matches the original cookie path
		HttpOnly: true,
		Secure:   false,                // Set to true if using HTTPS
		SameSite: http.SameSiteLaxMode, // Adjust SameSite as needed
	}

	http.SetCookie(w, &cookie) // Set the cookie to expire
	fmt.Println("Logout successful")
	// Optionally return a confirmation message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Logout successful"}`))
}

func GetAllInventaire(w http.ResponseWriter, r *http.Request) {
	inventaires, err := models.GetAllInventaire()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des inventaires", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(inventaires)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func GetAllPurchase(w http.ResponseWriter, r *http.Request) {
	achat, err := models.GetAllPurchase()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des achat", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(achat)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func GetAllMateriel(w http.ResponseWriter, r *http.Request) {
	materiel, err := models.GetAllMateriel()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des inventaires", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(materiel)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	employee, err := models.GetAllEmployee()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des inventaires", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(employee)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func GetListUser(w http.ResponseWriter, r *http.Request) {
	user, err := models.GetListUser()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des inventaires", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// getUserInfo handles the request and returns user information in JSON format
func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	// Retrieve the cookie named "jwt"
	cookie, err := r.Cookie("jwt")
	if err != nil {
		http.Error(w, "No JWT cookie found", http.StatusUnauthorized)
		return
	}

	// Parse the JWT token using the claims
	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil // Replace with your actual secret key
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Extract claims
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		http.Error(w, "Could not parse claims", http.StatusUnauthorized)
		return
	}

	// Fetch user from the database based on the claim issuer (user ID)
	var user models.User
	db := config.GetDb()
	err = db.Where("id = ?", claims.Issuer).First(&user).Error
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Marshal user to JSON
	res, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error marshalling user data", http.StatusInternalServerError)
		return
	}

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
