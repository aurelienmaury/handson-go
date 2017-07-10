package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"handson-go/cmn"
)

func main() {
	fmt.Println("Hello World")

	var gstate Cmn.GameState
	gstate.Partie = Cmn.WAIT
	gstate.Turn = "You"

	gstate.Grid = [3]map[string]string{
		map[string]string{"0": "-", "1":"-", "2":"-"},
		map[string]string{"0": "-", "1":"-", "2":"-"},
		map[string]string{"0": "-", "1":"-", "2":"-"},
	}


	http.HandleFunc("/game/state", func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode(gstate); err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", nil)

}
