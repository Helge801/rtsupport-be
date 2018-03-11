package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

//Handler for routerdirecting Router.Handler to proper function
type Handler func(*Client, interface{})

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// Router for
type Router struct {
	rules map[string]Handler
}

// NewRouter for creating new Router struct
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]Handler),
	}
}

// Handle for directing to proper function using handler
func (r *Router) Handle(msgName string, handler Handler) {
	r.rules[msgName] = handler
}

// ServeHTTP for
func (e *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
}
