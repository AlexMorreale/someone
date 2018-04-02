package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

var currentPostId int

func RedisConnect() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	HandleError(err)
	return c
}

func CreateQuote(q Quote) {
	currentPostId += 1

	q.Id = currentPostId
	q.Date = time.Now()

	c := RedisConnect()
	defer c.Close()

	b, err := json.Marshal(q)
	reply, err := c.Do("SET", "quote:"+strconv.Itoa(q.Id), b)
	HandleError(err)

	fmt.Println("GET ", reply)
}

func FindAll() Quotes {
	var quotes Quotes

	c := RedisConnect()
	defer c.Close()

	keys, err := c.Do("KEYS", "quote:*")
	HandleError(err)

	for _, k := range keys.([]interface{}) {
		var quote Quote

		reply, err := c.Do("GET", k.([]byte))
		HandleError(err)

		if err := json.Unmarshal(reply.([]byte), &quote); err != nil {
			panic(err)
		}

		quotes = append(quotes, quote)
	}
	return quotes
}

func FindRandom() Quote {
	var quote Quote

	c := RedisConnect()
	defer c.Close()

	key, err := c.Do("RANDOMKEY")
	HandleError(err)

	reply, err := c.Do("GET", key.([]byte))
	HandleError(err)

	if err := json.Unmarshal(reply.([]byte), &quote); err != nil {
		panic(err)
	}

	return quote
}
