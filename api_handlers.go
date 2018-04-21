package main

import (
	"encoding/json"
	mux "github.com/julienschmidt/httprouter"
	"net/http"
)

func ApiQuoteAll(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	var quotes Quotes
	quotes = FindAll() // We'll work on this
	HandleError(json.NewEncoder(w).Encode(quotes))
}

func ApiQuoteRandom(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	var quote Quote
	quote = FindRandom() // We'll work on this
	HandleError(json.NewEncoder(w).Encode(quote))
}
