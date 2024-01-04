package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/exgolden/BiblioGo/db"
	"github.com/exgolden/BiblioGo/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.Users
	db.DB.Find(&users)
	if len(users) == 0 {
		log.Println("No users found")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No users found"))
		return
	}
	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.Users
	db.DB.First(&user, params["boleta"])
	if user.ID == 0 {
		log.Println("User " + params["boleta"] + " not found")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	json.NewEncoder(w).Encode(&user)
}

func CreateUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user, userCheck models.Users
	json.NewDecoder(r.Body).Decode(&user)
	db.DB.Where("boleta = ?", user.Boleta).First(&userCheck)
	if userCheck.ID != 0 {
		log.Println("User " + user.Boleta + " already created")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User already created"))
		return
	}
	createUser := db.DB.Create(&user)
	err := createUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User not created"))
		return
	}
	log.Println("User " + user.Boleta + " not found while creating")
	log.Println("User " + user.Boleta + " created")
	json.NewEncoder(w).Encode(&user)
}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.Users
	db.DB.First(&user, params["boleta"])
	if user.ID == 0 {
		log.Println("User " + params["boleta"] + " not found while deleting")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)
}
