package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=isaibb password=Constructora4 dbname=bibliodb port=5432"
var DB *gorm.DB

func DBConnection() {
	var Derr error
	DB, Derr = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if Derr != nil {
		log.Println("Database connection error")
	} else {
		log.Println("Database connection succesfull")
	}

}
