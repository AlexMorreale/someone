package main

import (
	"fmt"
	mux "github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("tmpl/*"))

func renderTemplate(w http.ResponseWriter, tmpl string, q Quotes) {
	err := templates.ExecuteTemplate(w, tmpl+".html", q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Index(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func QuoteAll(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	var quotes Quotes
	quotes = FindAll()

	renderTemplate(w, "all", quotes)

}

func QuoteRandom(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	var quote Quote
	var quotes Quotes
	quote = FindRandom()
	quotes = append(quotes, quote)
	renderTemplate(w, "all", quotes)
}
