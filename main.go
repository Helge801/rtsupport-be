package main

import (
	"net/http"
)

//Channel for incoming WebSocket Channel
type Channel struct {
	ID   string      `json:"id"`
	Name interface{} `json:"name"`
}

func main() {
	router := NewRouter()

	router.Handle("channel add", addChannel)

	http.Handle("/", router)
	http.ListenAndServe(":4000", nil)
}
