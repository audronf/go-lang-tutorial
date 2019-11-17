package models

import (
	u "go-lang-tutorial/utils"
	"github.com/jinzhu/gorm"
	"reflect"
)

// Tweet struct
type Tweet struct {
	gorm.Model
	Text    string `json:"text"`
	UserID    uint `json:"userId"`
}

// NewTweet Create.
func (tweet *Tweet) NewTweet() map[string] interface{ } {
	var getTweet Tweet
	GetDB().Where("text = ? and user_id = ?", tweet.Text, tweet.UserID).First(&getTweet)
	if (reflect.DeepEqual(getTweet, Tweet{})) {
		GetDB().Create(tweet)
		response := u.Message(true, "Tweet sent succesfully")
		response["tweet"] = tweet
		return response
	}
	response := u.Message(false, "Don't send duplicated tweets")
	return response
}

// GetTweet by ID
func GetTweet(ID uint64) (*Tweet){
	var getTweet Tweet
	GetDB().Where("ID = ?", ID).Find(&getTweet)
	return &getTweet
}

// DeleteTweet D
func DeleteTweet(ID uint64) map[string] interface{ } {
	var tweet Tweet
	GetDB().Where("id = ?", ID).Delete(tweet)
	response := u.Message(true, "Tweet deleted succesfully")
	return response
}
