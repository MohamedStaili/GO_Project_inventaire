package controllers

import (
	"MohamedStaili/GO_Project_inventaire/pkg/models"
	utils "MohamedStaili/GO_Project_inventaire/pkg/utiles"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

func SupprimerInventaire(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
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

// Handler HTTP pour modifier un inventaire
func ModifierInventaire(w http.ResponseWriter, r *http.Request) {
	// Créer une instance de modèle Inventaire pour stocker les données de la requête
	modInv := models.Inventaire{}

	// Parser le corps de la requête JSON dans modInv
	utils.ParseBody(r, &modInv)

	// Récupérer l'ID de l'inventaire à partir des variables de requête
	/*vars := mux.Vars(r)
	invId := vars["id"]
	Id, err := strconv.ParseInt(invId, 10, 64) // 10 base, 64 bits
	if err != nil {
		http.Error(w, "Erreur de parsing de l'ID", http.StatusBadRequest)
		return
	}

	// Affecter l'ID à modInv.ID
	modInv.ID = uint(Id)*/

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
