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

	router.GET("/api/v1/leaderboard/:key", handlers.GetLeaderboard)
	router.GET("/api/v1/user/:username/statistics", handlers.GetUserStats)
	router.GET("/api/v1/user/:username/friends", handlers.GetUserFriends)

	log.Fatal(http.ListenAndServe(":8080", router))
}
