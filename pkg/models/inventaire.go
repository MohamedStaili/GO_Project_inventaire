package models

import (
	"MohamedStaili/GO_Project_inventaire/pkg/config"

	"github.com/jinzhu/gorm"
)

// Déclaration de la variable de la base de données
var db *gorm.DB

// Structure Inventaire
type Inventaire struct {
	gorm.Model        // Embedding gorm.Model pour inclure des champs comme ID, CreatedAt, UpdatedAt, DeletedAt
	NInv       string `json:"n_inv"`      // Numéro d'inventaire
	IdMat      int    `json:"id_mat"`     // ID du matériel
	NumeroSer  string `json:"numero_ser"` // Numéro de série
	RAchat     string `json:"r_achat"`    // Référence d'achat
	IdEmploye  int    `json:"id_employe"` // ID de l'employé
}

// Structure Employe
type Employe struct {
	gorm.Model        // Embedding gorm.Model pour inclure des champs comme ID, CreatedAt, UpdatedAt, DeletedAt
	IdEmploye  int    `json:"id_employe"`
	Nom        string `json:"nom"`
	Prenom     string `json:"prenom"`
}

// Structure Materiel
type Materiel struct {
	gorm.Model         // Embedding gorm.Model pour inclure des champs comme ID, CreatedAt, UpdatedAt, DeletedAt
	IdMat       int    `json:"id_mat"`
	MatLabel    string `json:"mat_label"`
	MarqueModel string `json:"marque_model"`
}

// Structure Achat
type Achat struct {
	gorm.Model          // Embedding gorm.Model pour inclure des champs comme ID, CreatedAt, UpdatedAt, DeletedAt
	RefAchat    string  `json:"ref_achat"`
	NumFact     string  `json:"num_fact"`
	PriAchatHT  float64 `json:"pri_achat_ht"`
	Fournisseur string  `json:"fournisseur"`
	DateEntree  string  `json:"date_entree"`
	IdMat       string  `json:"id_mat"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Inventaire{}, &Employe{}, &Materiel{}, &Achat{})

}
