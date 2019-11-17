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

// GetTweet by ID
func GetTweet(ID uint64) (*Tweet , *gorm.DB){
	var getTweet Tweet
	db:=db.Where("ID = ?", ID).Find(&getTweet)
	return &getTweet, db
}

// DeleteTweet D
func DeleteTweet(ID uint64) map[string] interface{ } {
	// GetDB().Exec("DELETE FROM Tweets WHERE id = $1", ID)
	var tweet Tweet
	GetDB().Where("id = ?", ID).Delete(tweet)
	response := u.Message(true, "Tweet deleted succesfully")
	return response
}
