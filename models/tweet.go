package models

import (
	u "go-lang-tutorial/utils"
	"github.com/jinzhu/gorm"
)

// Tweet struct
type Tweet struct {
	gorm.Model
	Text    string `json:"text"`
	User    User `json:"user"`
}

// Create tweet
func (tweet *Tweet) NewTweet() map[string]interface{} {
	db.Create(tweet)

	response := u.Message(true, "Account has been created")
	response["tweet"] = tweet
	return response
}
