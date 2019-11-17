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

// NewTweet Create
func (tweet *Tweet) NewTweet() map[string] interface{ } {
	GetDB().Create(tweet)

	response := u.Message(true, "Tweet sent succesfully")
	response["tweet"] = tweet
	return response
}

// DeleteTweet D
func DeleteTweet(ID uint64) map[string] interface{ } {
	GetDB().Exec("DELETE FROM Tweets WHERE id = $1", ID)
	response := u.Message(true, "Tweet deleted succesfully")
	return response
}
