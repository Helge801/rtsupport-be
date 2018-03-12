package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// Message for
type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

//Client for
type Client struct {
	send   chan Message
	socket *websocket.Conn
}

func (c *Client) Write() {
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err != nil {
			fmt.Println(err)
			break
		}
	}
	c.socket.Close()
}

//NewClient return new Client Struct
func NewClient(socket *websocket.Conn) *Client {
	return &Client{
		send:   make(chan Message),
		socket: socket,
	}
}
