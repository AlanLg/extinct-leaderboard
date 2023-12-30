package handlers

import (
	"encoding/json"
	"first-web-app/database"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Leaderboard(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	users := database.GetLeaderboard(p.ByName("key"))

	for _, user := range users {
		output, err := json.MarshalIndent(user, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}

	json, err := json.Marshal(users)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("users could not be converted to json"))
		panic("could not convers users to json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
