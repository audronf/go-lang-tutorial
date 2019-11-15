package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"./models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func main() {
	router := mux.NewRouter()
	dbURI := fmt.Sprintf("postgres://dev:dev@localhost:5432/golangtutorial?sslmode=disable")
	db, err = gorm.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&models.User{})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
