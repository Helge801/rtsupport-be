package main

import (
	"net/http"
)

//Message for incoming WebSocket Message
type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

//Channel for incoming WebSocket Channel
type Channel struct {
	ID   string      `json:"id"`
	Name interface{} `json:"name"`
}

func main() {
	router := &Router{}

	router.Handle("channel add", addChannel)

	http.Handle("/", router)
	http.ListenAndServe(":4000", nil)
}
