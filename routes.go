package main

import (
	mux "github.com/julienschmidt/httprouter"
)

type Route struct {
	Method  string
	Pattern string
	Handle  mux.Handle //httprouter package as mux
}

type Routes []Route

var routes = Routes{
	Route{
		"GET",
		"/",
		Index,
	},
	Route{
		"GET",
		"/api/all",
		ApiQuoteAll,
	},
	Route{
		"GET",
		"/api/random",
		ApiQuoteRandom,
	},
	Route{
		"GET",
		"/all",
		QuoteAll,
	},
	Route{
		"GET",
		"/random",
		QuoteRandom,
	},
}
