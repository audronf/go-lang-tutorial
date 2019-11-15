package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {
	dBURI:= fmt.Sprintf("postgres://dev:dev@localhost:5432/golangtutorial?sslmode=disable")
	fmt.Println(dBURI)

	conn, err := gorm.Open("postgres", dBURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&User{})
}

func GetDB() *gorm.DB {
	return db
}
