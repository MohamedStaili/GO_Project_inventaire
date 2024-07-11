package models

import (
	"MohamedStaili/GO_Project_inventaire/pkg/config"
	"fmt"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Structure Inventaire
type Inventaire struct {
	gorm.Model
	DateAffectation string   `json:"date_affectation"`
	NInv            string   `json:"n_inv"`
	IdMat           uint     `json:"id_mat"`
	Materiel        Materiel `gorm:"foreignkey:IdMat" json:"materiel"`
	NumeroSer       string   `json:"numero_ser"`
	IdAchat         uint     `json:"id_achat"`
	Achat           Achat    `gorm:"foreignkey:IdAchat" json:"achat"`
	IdEmploye       uint     `json:"id_employe"`
	Employe         Employe  `gorm:"foreignkey:IdEmploye" json:"employe"`
}

// Structure Achat
type Achat struct {
	gorm.Model
	RefAchat    string  `json:"ref_achat"`
	NumFact     string  `json:"num_fact"`
	PriAchatHT  float64 `json:"pri_achat_ht"`
	Fournisseur string  `json:"fournisseur"`
	DateEntree  string  `json:"date_entree"`
	IdMat       uint    `json:"id_mat"`
}

// Structure Employe
type Employe struct {
	gorm.Model
	Nom    string `json:"nom"`
	Prenom string `json:"prenom"`
}

// Structure Materiel
type Materiel struct {
	gorm.Model
	MatLabel    string `json:"mat_label"`
	MarqueModel string `json:"marque_model"`
}

// Nommer explicitement la table
func (Inventaire) TableName() string {
	return "inventaires"
}

func (Employe) TableName() string {
	return "employes"
}

func (Achat) TableName() string {
	return "achats"
}

func (Materiel) TableName() string {
	return "materiels"
}

// Inside your init() function or database initialization code
func init() {
	config.Connect()
	db = config.GetDb()
	db.LogMode(true)
	db.AutoMigrate(&Materiel{}, &Achat{}, &Employe{}, &Inventaire{})

	// Auto-migration avec les relations et les clés étrangères
	db.Table("inventaires").AutoMigrate(&Inventaire{}).
		AddForeignKey("id_mat", "materiels(id)", "CASCADE", "CASCADE").
		AddForeignKey("id_achat", "achats(id)", "CASCADE", "CASCADE").
		AddForeignKey("id_employe", "employes(id)", "CASCADE", "CASCADE")
	db.Table("achats").AutoMigrate(&Achat{}).AddForeignKey("id_mat", "materiels(id)", "CASCADE", "CASCADE")
}

func (b *Inventaire) AjouterInventaire() (*Inventaire, error) {
	// Vérifiez que les clés étrangères existent avant l'insertion
	if err := db.Where("id = ?", b.IdMat).First(&Materiel{}).Error; err != nil {
		return nil, fmt.Errorf("Materiel with id %d does not exist", b.IdMat)
	}
	if err := db.Where("id = ?", b.IdAchat).First(&Achat{}).Error; err != nil {
		return nil, fmt.Errorf("Achat with id %d does not exist", b.IdAchat)
	}
	if err := db.Where("id = ?", b.IdEmploye).First(&Employe{}).Error; err != nil {
		return nil, fmt.Errorf("Employe with id %d does not exist", b.IdEmploye)
	}

	// Ajouter l'inventaire si toutes les clés étrangères sont valides
	db.NewRecord(&b)
	db.Create(&b)
	return b, nil
}

func GetALLInventaire() []Inventaire {
	var inv []Inventaire
	db.Preload("Achat").Preload("Employe").Preload("Materiel").Find(&inv)
	return inv
}

func SearchPage(Id int64) (Inventaire, *gorm.DB) {
	var getInv Inventaire
	db.Preload("Achat").Preload("Employe").Preload("Materiel").Where("id = ?", Id).Find(&getInv)
	return getInv, db
}

func SupprimerInventaire(Id int64) error {
	var delInv Inventaire
	// Log de débogage pour vérifier l'ID
	fmt.Printf("Attempting to delete record with ID: %d\n", Id)

	// Trouvez l'inventaire à supprimer
	if err := db.Where("id = ?", Id).First(&delInv).Error; err != nil {
		fmt.Printf("Error finding record: %v\n", err)
		return err
	}

	// Log de débogage pour vérifier l'enregistrement trouvé
	fmt.Printf("Record found: %+v\n", delInv)

	// Supprimez l'inventaire
	if err := db.Delete(&delInv).Error; err != nil {
		fmt.Printf("Error deleting record: %v\n", err)
		return err
	}

	fmt.Println("Record deleted successfully")
	return nil
}

func (b *Inventaire) ModifierInventaire() *Inventaire {
	if db.NewRecord(&b) {
		fmt.Println("Erreur: Impossible de mettre à jour, l'inventaire n'existe pas encore.")
		return b
	}
	db.Save(&b)
	return b
}
