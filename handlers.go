package main

import (
	"encoding/json"
	"fmt"
	mux "github.com/julienschmidt/httprouter"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func QuoteAll(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	var quotes Quotes
	quotes = FindAll() // We'll work on this
	HandleError(json.NewEncoder(w).Encode(quotes))
}

func QuoteRandom(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	var quote Quote
	quote = FindRandom() // We'll work on this
	HandleError(json.NewEncoder(w).Encode(quote))
}
