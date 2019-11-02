package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
	"go-lang-tutorial/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func main() {
	router := mux.NewRouter()
	dbURI := fmt.Sprintf("postgres://postgres:%s@%s:5432/%s?sslmode=disable",
						os.Getenv("db_pass"), os.Getenv("db_host"), os.Getenv("db_name"))
	db, err = gorm.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&models.User{})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
