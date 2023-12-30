package main

import (
	"extinct-leaderboard/database"
	"extinct-leaderboard/handlers"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	database.Init("mongodb://localhost:27017")

	router := httprouter.New()
	router.GET("/leaderboard/:key", handlers.Leaderboard)

	log.Fatal(http.ListenAndServe(":8080", router))
}
