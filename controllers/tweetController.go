package controllers

import (
	"go-lang-tutorial/models"
	u "go-lang-tutorial/utils"
	"encoding/json"
	"net/http"
)

// NewTweet - Sends a new tweet if the user is logged in
var NewTweet = func(w http.ResponseWriter, r *http.Request) {

	tweet := &models.Tweet{}
	user := r.Context().Value("user").(uint)
	err := json.NewDecoder(r.Body).Decode(tweet) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	tweet.UserID = user
	resp := tweet.NewTweet() //Create tweet
	u.Respond(w, resp)
}
