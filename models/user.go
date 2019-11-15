package models

import (
	"github.com/dgrijalva/jwt-go"
	u "go-lang-tutorial/utils"
	"github.com/jinzhu/gorm"
	"os"
	"golang.org/x/crypto/bcrypt"
)

// Token model
type Token struct {
	UserID uint
	jwt.StandardClaims
}

// User model
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

// Create account
func (account *User) Create() map[string]interface{} {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)
	
	db.Create(account)
	//Create new JWT token for the newly registered account
	tk := &Token{UserID: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" //delete password

	response := u.Message(true, "Account has been created")
	response["account"] = account
	
	return response
}
