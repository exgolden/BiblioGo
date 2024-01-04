package main

import (
	"log"
	"net/http"

	"github.com/exgolden/BiblioGo/db"
	"github.com/exgolden/BiblioGo/models"
	"github.com/exgolden/BiblioGo/routes"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server up at port 3000")
	// Modulo de biblioteca
	db.DBConnection()
	db.DB.AutoMigrate(models.Books{})
	db.DB.AutoMigrate(models.Users{})

	r := mux.NewRouter()

	// Rutas de usuarios
	r.HandleFunc("/users/", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{boleta}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/create", routes.CreateUsersHandler).Methods("POST")
	r.HandleFunc("/users/delete/{boleta}", routes.DeleteUsersHandler).Methods("Delete")

	http.ListenAndServe(":3000", r)

}
