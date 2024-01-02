package handlers

import (
	"encoding/json"
	"extinct-leaderboard/database"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetLeaderboard(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	users := database.GetLeaderboard(p.ByName("key"))
	json, err := json.Marshal(users)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server Error"))
		panic("could not convers users to json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func GetUserStats(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user, err := database.GetUser(p.ByName("username"))
	json, jsonErr := json.Marshal(user.StatisticsUser)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server Error"))
		panic("user with ign " + p.ByName("username") + " not found")
	}

	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server Error"))
		panic("could not convers users to json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func GetUserFriends(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
