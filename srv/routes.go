package main

import (
	"net/http"
	 "github.com/aurelienmaury/handson-go/cmn"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/game/state/{joueur}",
		GetStatePlayer,
	},
}