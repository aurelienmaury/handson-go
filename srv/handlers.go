package main

import (
	"net/http"
	"encoding/json"
	"github.com/aurelienmaury/handson-go/cmn"
)

func GetStatePlayer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(); err != nil {
		panic(err)
	}
}