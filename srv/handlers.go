package main

import (
	"net/http"
	"encoding/json"
	"handson-go/cmn"

)

func GetStatePlayer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var gstate Cmn.GameState
	gstate.Partie = Cmn.WAIT
	gstate.Turn = "You"

	if err := json.NewEncoder(w).Encode(gstate); err != nil {
		panic(err)
	}
}