package controllers

import (
	"github.com/gorilla/mux"
	"go-lang-tutorial/models"
	u "go-lang-tutorial/utils"
	"strconv"
	"encoding/json"
	"net/http"
	"fmt"
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

// UpdateTweet U
func UpdateTweet(w http.ResponseWriter, r *http.Request) {
	var updateTweet = &models.Tweet{}
	user := r.Context().Value("user").(uint)
	u.ParseBody(r, updateTweet)
	vars := mux.Vars(r)
	tweetID := vars["tweetID"]
	ID, err:= strconv.ParseUint(tweetID, 0, 0)
	if err != nil {
		fmt.Println(err.Error())
	}
	tweetDetails := models.GetTweet(ID)
	if (tweetDetails.UserID != user) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if updateTweet.Text != "" {
		tweetDetails.Text = updateTweet.Text
	}

	models.GetDB().Save(&tweetDetails)
	resp, _ := json.Marshal(tweetDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// DeleteTweet - Deletes a tweet
var DeleteTweet = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tweetID := vars["tweetID"]
	user := r.Context().Value("user").(uint)
	ID, err := strconv.ParseUint(tweetID, 0, 0)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}
	tweetDetails := models.GetTweet(ID)
	if (tweetDetails.UserID != user) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp := models.DeleteTweet(ID)
	u.Respond(w, resp)
}
