package models

import (
	u "go-lang-tutorial/utils"
	"github.com/jinzhu/gorm"
)

// Tweet struct
type Tweet struct {
	gorm.Model
	Text    string `json:"text"`
	UserID    uint `json:"userId"`
}

// Create tweet
func (tweet *Tweet) NewTweet() map[string]interface{} {
	GetDB().Create(tweet)

	response := u.Message(true, "Tweet sent succesfully")
	response["tweet"] = tweet
	return response
}
