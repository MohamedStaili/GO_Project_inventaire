package models

import (
	"MohamedStaili/GO_Project_inventaire/pkg/config"
	"errors"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt"
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
	DateAffectation string `json:"Assignement Date"`
	NInv            string `json:"Inventory Number"`
	IdMat           uint   `json:"Materiel ID"`
	//Materiel        Materiel `gorm:"foreignkey:IdMat" json:"materiel"`
	NumeroSer string `json:"Serie Number"`
	IdAchat   uint   `json:"Purchase ID"`
	//Achat           Achat    `gorm:"foreignkey:IdAchat" json:"achat"`
	IdEmploye uint `json:"Employee ID"`
	//Employe         Employe  `gorm:"foreignkey:IdEmploye" json:"employe"`
}
type CustomClaims struct {
	ID   uint   `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

// Structure Achat
type Achat struct {
	gorm.Model
	RefAchat    string  `json:"Purchase Ref"`
	NumFact     string  `json:"Facteur Number"`
	PriAchatHT  float64 `json:"Purchase Price H.T"`
	Fournisseur string  `json:"Fournisseur"`
	DateEntree  string  `json:"Entre Date"`
	IdMat       uint    `json:"Materiel ID"`
	//Materiel    Materiel `gorm:"foreignkey:IdMat" json:"materiel"`
}

// Structure Employe
type Employe struct {
	gorm.Model
	Nom    string `json:"Last Name"`
	Prenom string `json:"Firs Name"`
}

// Structure Materiel
type Materiel struct {
	gorm.Model
	MatLabel    string `json:"Materiel Lbale"`
	MarqueModel string `json:"Marque/Model"`
}

// Structure User
type User struct {
	gorm.Model
	UUID     string `json:"uuid"`
	Username string `json:"Username"`
	Email    string `json:"Email" `
	Password string `json:"Password" ` //pour ne pas returner le password
	Role     string `json:"Role"`
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
	db.Table("achats").AutoMigrate(&Achat{}).
		AddForeignKey("id_mat", "materiels(id)", "CASCADE", "CASCADE")
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
func (b *Achat) AjouterAchat() (*Achat, error) {
	// Vérifiez que les clés étrangères existent avant l'insertion
	if err := db.Where("id = ?", b.IdMat).First(&Materiel{}).Error; err != nil {
		return nil, fmt.Errorf("Materiel with id %d does not exist", b.IdMat)
	}

	// Ajouter l'inventaire si toutes les clés étrangères sont valides
	db.NewRecord(&b)
	db.Create(&b)
	return b, nil
}
func (b *Employe) AjouterEmployee() (*Employe, error) {

	db.NewRecord(&b)
	db.Create(&b)
	return b, nil
}
func (b *Materiel) AjouterMateriel() (*Materiel, error) {

	db.NewRecord(&b)
	db.Create(&b)
	return b, nil
}

func GetAllInventaire() ([]Inventaire, error) {
	var inventaires []Inventaire
	err := db.Find(&inventaires).Error
	if err != nil {
		return nil, err
	}
	return inventaires, nil
}
func GetAllMateriel() ([]Materiel, error) {
	var materiel []Materiel
	err := db.Find(&materiel).Error
	if err != nil {
		return nil, err
	}
	return materiel, nil
}
func GetAllEmployee() ([]Employe, error) {
	var employe []Employe
	err := db.Find(&employe).Error
	if err != nil {
		return nil, err
	}
	return employe, nil
}
func GetAllPurchase() ([]Achat, error) {
	var achat []Achat
	err := db.Find(&achat).Error
	if err != nil {
		return nil, err
	}
	return achat, nil
}

func GetListUser() ([]User, error) {
	var user []User
	err := db.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
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
func SupprimerMateriel(Id int64) error {
	var delMat Materiel
	// Log de débogage pour vérifier l'ID
	fmt.Printf("Attempting to delete record with ID: %d\n", Id)

	// Trouvez l'inventaire à supprimer
	if err := db.Where("id = ?", Id).First(&delMat).Error; err != nil {
		fmt.Printf("Error finding record: %v\n", err)
		return err
	}

	// Log de débogage pour vérifier l'enregistrement trouvé
	fmt.Printf("Record found: %+v\n", delMat)

	// Supprimez l'inventaire
	if err := db.Delete(&delMat).Error; err != nil {
		fmt.Printf("Error deleting record: %v\n", err)
		return err
	}

	fmt.Println("Record deleted successfully")
	return nil
}
func SupprimerEmploye(Id int64) error {
	var delEmp Employe
	// Log de débogage pour vérifier l'ID
	fmt.Printf("Attempting to delete record with ID: %d\n", Id)

	// Trouvez l'Employe à supprimer
	if err := db.Where("id = ?", Id).First(&delEmp).Error; err != nil {
		fmt.Printf("Error finding record: %v\n", err)
		return err
	}

	// Log de débogage pour vérifier l'enregistrement trouvé
	fmt.Printf("Record found: %+v\n", delEmp)

	// Supprimez l'inventaire
	if err := db.Delete(&delEmp).Error; err != nil {
		fmt.Printf("Error deleting record: %v\n", err)
		return err
	}

	fmt.Println("Record deleted successfully")
	return nil
}
func SupprimerAchat(Id int64) error {
	var delAchat Achat
	// Log de débogage pour vérifier l'ID
	fmt.Printf("Attempting to delete record with ID: %d\n", Id)

	// Trouvez l'Achat à supprimer
	if err := db.Where("id = ?", Id).First(&delAchat).Error; err != nil {
		fmt.Printf("Error finding record: %v\n", err)
		return err
	}

	// Log de débogage pour vérifier l'enregistrement trouvé
	fmt.Printf("Record found: %+v\n", delAchat)

	// Supprimez l'Achat
	if err := db.Delete(&delAchat).Error; err != nil {
		fmt.Printf("Error deleting record: %v\n", err)
		return err
	}

	fmt.Println("Record deleted successfully")
	return nil
}

func (b *Inventaire) ModifierInventaire() (*Inventaire, error) {
	var existingMateriel Inventaire
	// Recherchez le matériel par son ID
	if err := db.Where("id = ?", b.ID).First(&existingMateriel).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println("Erreur: INVENTAIRE n'existe pas.")
			return nil, err
		}
		return nil, err
	}

	// Si le matériel existe, effectuez la mise à jour
	if err := db.Save(b).Error; err != nil {
		return nil, err
	}

	return b, nil
}

func (b *Materiel) ModifierMateriel() (*Materiel, error) {
	var existingMateriel Materiel
	// Recherchez le matériel par son ID
	if err := db.Where("id = ?", b.ID).First(&existingMateriel).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println("Erreur: Materiel n'existe pas.")
			return nil, err
		}
		return nil, err
	}

	// Si le matériel existe, effectuez la mise à jour
	if err := db.Save(b).Error; err != nil {
		return nil, err
	}

	return b, nil
}
func (b *Employe) ModifierEmploye() (*Employe, error) {
	var existingMateriel Employe
	// Recherchez le matériel par son ID
	if err := db.Where("id = ?", b.ID).First(&existingMateriel).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println("Erreur: Employe n'existe pas.")
			return nil, err
		}
		return nil, err
	}

	// Si le matériel existe, effectuez la mise à jour
	if err := db.Save(b).Error; err != nil {
		return nil, err
	}

	return b, nil
}
func (b *Achat) ModifierAchat() (*Achat, error) {
	var existingMateriel Achat
	// Recherchez le matériel par son ID
	if err := db.Where("id = ?", b.ID).First(&existingMateriel).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println("Erreur: Employe n'existe pas.")
			return nil, err
		}
		return nil, err
	}

	// Si le matériel existe, effectuez la mise à jour
	if err := db.Save(b).Error; err != nil {
		return nil, err
	}

	return b, nil
}

func (b *User) ModifierProfile() (*User, error) {
	hashPassword(b.Password)
	var existingUser User
	// Recherchez le matériel par son ID
	if err := db.Where("id = ?", b.ID).First(&existingUser).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println("Erreur: User n'existe pas.")
			return nil, err
		}
		return nil, err
	}

	// Si le matériel existe, effectuez la mise à jour
	if err := db.Save(b).Error; err != nil {
		return nil, err
	}

	return b, nil
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
	/*if existingUser.Role == "admin" {
		return errors.New("invalid role")
	}*/
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

func sendWelcomeEmail(email string) error {

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
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func LoginPage(email, password string) (error, User) {
	var user User

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return errors.New("email does not exist"), user
	}

	if !checkPasswordHash(password, user.Password) {
		return errors.New("incorrect password"), user
	}

	fmt.Println("Login Successful")
	return nil, user
}

func UserByEmail(email string) (User, error) {
	var user User
	if err := db.Where("email=?", email).First(&user).Error; err != nil {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

/*func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		tmpl := template.Must(template.ParseFiles(page))
		layouts, err := filepath.Glob("./templates/layouts/*.layout.tmpl")
		if err != nil {
			return cache, err
		}
		if len(layouts) > 0 {
			tmpl.ParseGlob("./templates/layouts/*.layout.tmpl")
		}
		cache[name] = tmpl
	}
	return cache, nil
}
*/
