package main

import (
	"time"
)

type Quote struct {
	Id   int       `json:"id"`
	Text string    `json:"text"`
	Date time.Time `json:"date"`
}

type Quotes []Quote
