package controllers

import (
	"go-lang-tutorial/models"
	u "go-lang-tutorial/utils"
	"net/http"
)

var NewTweet = func(w http.ResponseWriter, r *http.Request) {

	tweet := &models.Tweet{}
	resp := tweet.NewTweet() //Create tweet
	u.Respond(w, resp)
}
