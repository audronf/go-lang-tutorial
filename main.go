package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-lang-tutorial/app"
	"go-lang-tutorial/controllers"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/tweet/new", controllers.NewTweet).Methods("POST")
	router.HandleFunc("/api/tweet/{tweetID}", controllers.DeleteTweet).Methods("DELETE")
	router.HandleFunc("/api/tweet/{tweetID}", controllers.UpdateTweet).Methods("PUT")
	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
