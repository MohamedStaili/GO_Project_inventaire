package models

import (
	"MohamedStaili/GO_Project_inventaire/pkg/config"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/gomail.v2"
)

var db *gorm.DB
var validate *validator.Validate

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

// Structure User
type User struct {
	gorm.Model
	UUID     string `json:"uuid"`
	Username string `json:"username" validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role" gorm:"default:user"`
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
func (User) TableName() string {
	return "user"
}

// Inside your init() function or database initialization code
func init() {
	config.Connect()
	db = config.GetDb()
	db.LogMode(true)
	db.AutoMigrate(&Materiel{}, &Achat{}, &Employe{}, &Inventaire{}, &User{})

	// Auto-migration avec les relations et les clés étrangères
	db.Table("inventaires").AutoMigrate(&Inventaire{}).
		AddForeignKey("id_mat", "materiels(id)", "CASCADE", "CASCADE").
		AddForeignKey("id_achat", "achats(id)", "CASCADE", "CASCADE").
		AddForeignKey("id_employe", "employes(id)", "CASCADE", "CASCADE")
	db.Table("achats").AutoMigrate(&Achat{}).AddForeignKey("id_mat", "materiels(id)", "CASCADE", "CASCADE")
	validate = validator.New()
}
func (u *User) Validate() error {
	return validate.Struct(u)
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

// Create a user
func (U *User) AjouterUser() (*User, error) {
	if err := U.Validate(); err != nil {
		return nil, err
	}
	db.NewRecord(U)
	db.Create(U)
	return U, nil
}
func SupprimerUser(Id int64) error {
	var delUser User
	if err := db.Where("id=?", Id).First(&delUser).Error; err != nil {
		fmt.Printf("Error finding User:%v", err)
		return err
	}
	if err := db.Unscoped().Delete(&delUser).Error; err != nil { //Delete permanently:Unscoped()
		fmt.Printf("Error User not deleted:%v", err)
		return err
	}
	fmt.Println("User deleted succefly")
	return nil
}

// GORM hook: BeforeCreate
func (U *User) BeforeCreate(scope *gorm.Scope) (err error) {
	U.UUID = uuid.New().String()
	// Vérifier si l'email existe déjà avant la création
	var existingUser User
	if err := db.Where("email = ?", U.Email).First(&existingUser).Error; err == nil {
		return errors.New("email already exists")
	}
	// Hash the password before saving
	if U.Password, err = hashPassword(U.Password); err != nil {
		return err
	}

	// Ensure the role is not "admin"
	if existingUser.UUID.Role == "admin" {
		return errors.New("invalid role")
	}
	return nil
}

// GORM hook: AfterCreate
func (U *User) AfterCreate(scope *gorm.Scope) error {
	// Envoyer un email de bienvenue à l'utilisateur
	go func() {
		if err := sendWelcomeEmail(U.Email); err != nil {
			log.Printf("Failed to send welcome email to %s: %v\n", U.Email, err)
		}
	}()
	return nil
}

// Fonction pour envoyer un email de bienvenue (exemple simplifié)
func sendWelcomeEmail(email string) error {
	// Ici, vous pouvez implémenter la logique d'envoi d'email
	// Utilisation de bibliothèques comme gomail.v2
	// Exemple simplifié :
	m := gomail.NewMessage()
	m.SetHeader("From", "inventaire_contact@devoteam.ac.ma")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Welcome!")
	m.SetBody("text/plain", "Welcome to our platform!")
	d := gomail.NewDialer("smtp.gmail.com", 587, "5254mohamed@gmail.com", "btnjdvwekweycsss")
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

// Fonction pour hacher les mots de passe
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func LoginPage(email, password string) error {
	var user User
	if err := db.Where("email=?", email).First(&user).Error; err != nil {

		errors.New("email not exist")
		return err
	} else {
		if err := db.Where("password=?", password).First(&user).Error; err != nil {
			errors.New("incorrect password")
		} else {
			fmt.Println("Login Succeful")
			return nil
		}

	}
	return nil

}
