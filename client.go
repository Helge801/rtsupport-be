package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// FindHandler for
type FindHandler func(string) (Handler, bool)

// Message for
type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

//Client for
type Client struct {
	send        chan Message
	socket      *websocket.Conn
	findHandler FindHandler
}

func (c *Client) Read() {
	var message Message
	for {
		if err := c.socket.ReadJSON(&message); err != nil {
			break
		}
		if handler, found := c.findHandler(message.Name); found {
			handler(c, message.Data)
		}
	}
	c.socket.Close()
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
func NewClient(socket *websocket.Conn, findHandler FindHandler) *Client {
	return &Client{
		send:        make(chan Message),
		socket:      socket,
		findHandler: findHandler,
	}
}
