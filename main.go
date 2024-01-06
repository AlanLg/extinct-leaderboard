package main

import (
	"extinct-leaderboard/database"
	"extinct-leaderboard/handlers"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	database.Init(os.Getenv("DATABASE_URI"))

	router := httprouter.New()

	router.GET("/api/v1/leaderboard/:key", handlers.GetLeaderboard)
	router.GET("/api/v1/user/:username/statistics", handlers.GetUserStats)
	router.GET("/api/v1/user/:username/friends", handlers.GetUserFriends)

	log.Fatal(http.ListenAndServe(":8080", router))
}
