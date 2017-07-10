package main

import (
	"fmt"
)
import "net/http"
import "github.com/aurelienmaury/handson-go/cmn"

func main() {
	fmt.Println("Hello World")



	http.HandleFunc("/game/state/{joueur}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>A simple web server</h1>"))
	})

	http.ListenAndServe(":8080", nil)

}
