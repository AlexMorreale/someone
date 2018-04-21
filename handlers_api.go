package main

import (
	"encoding/json"
	mux "github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
)

func ApiQuoteAll(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	var quotes Quotes
	quotes = FindAll()
	HandleError(json.NewEncoder(w).Encode(quotes))
}

func ApiQuoteRandom(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	var quote Quote
	quote = FindRandom()
	HandleError(json.NewEncoder(w).Encode(quote))
}

func ApiQuoteCreate(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	var quote Quote

	// Read request body and close it
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	HandleError(err)
	defer r.Body.Close()
	// Save JSON to quote struct
	if err := json.Unmarshal(body, &quote); err != nil {
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	CreateQuote(quote)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}
