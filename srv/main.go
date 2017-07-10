package main

import (
	"fmt"
	"handson-go/cmn"
	"net/http"
	"encoding/json"
)

func main() {
	fmt.Println("Hello World")

	var gstate Cmn.GameState
	gstate.Partie = Cmn.WAIT
	gstate.Turn = "You"

	http.HandleFunc("/game/state", func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode(gstate); err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", nil)

}
