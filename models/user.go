package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

/*
JWT claims struct
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}

//a struct to rep user account
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}
