package main

import mux "github.com/julienschmidt/httprouter"

func NewRouter() *mux.Router {
	router := mux.New()
	for _, route := range routes {
		router.Handle(route.Method, route.Pattern, route.Handle)
	}
	return router
}
