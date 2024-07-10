package models

import (
	"MohamedStaili/GO_Project_inventaire/pkg/config"
	"fmt"

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
func (b *Inventaire) AjouterInventaire() *Inventaire { //Creer un nouveau inventaire
	db.NewRecord(&b)
	db.Create(&b)
	return b
}
func GetALLInventaire() []Inventaire { //pour recuperer toute les inventaires de la base de donnee
	var inv []Inventaire
	db.Find(&inv)
	return inv
}
func SearchPage(Id int64) (Inventaire, *gorm.DB) { //cette,fct est a developpe,incha'allah
	var getInv Inventaire
	db.Where("id=?", Id).Find(&getInv)
	return getInv, db
}
func SupprimerInventaire(Id int64) Inventaire {
	var delInv Inventaire
	db.Where("id=?", Id).Delete(delInv)
	return delInv
}
func (b *Inventaire) ModifierInventaire() *Inventaire {
	// Vérifie si b est un nouvel enregistrement
	if db.NewRecord(&b) {
		// Gère le cas où b est un nouvel enregistrement
		fmt.Println("Erreur: Impossible de mettre à jour, l'inventaire n'existe pas encore.")
		return b
	}

	// Met à jour l'inventaire dans la base de données
	db.Save(&b)

	return b
}
