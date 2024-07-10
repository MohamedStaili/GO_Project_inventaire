package controllers

import (
	"MohamedStaili/GO_Project_inventaire/pkg/models"
	utils "MohamedStaili/GO_Project_inventaire/pkg/utiles"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK) //code=200
	w.Write(res)
}
func AjouterInventaire(w http.ResponseWriter, r *http.Request) {
	NewInv := models.Inventaire{}
	utils.ParseBody(r, NewInv)
	b := NewInv.AjouterInventaire()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
