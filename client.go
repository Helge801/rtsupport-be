package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Message for
type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

//Client for
type Client struct {
	send chan Message
}

func (c *Client) write() {
	for msg := range c.send {
		//TODO: socket.sendJSON(msg)
		fmt.Printf("%#v\n", msg)
	}
}

func (c *Client) subscribeChannels() {
	// TODO: changefeed Query RethinkDB
	for {
		time.Sleep(r())
		c.send <- Message{"channel add", ""}
	}
}

func (c *Client) subscribeMessages() {
	// TODO: changefeed Query RethinkDB
	for {
		time.Sleep(r())
		c.send <- Message{"message add", ""}
	}
}

func r() time.Duration {
	return time.Millisecond * time.Duration(rand.Intn(1000))
}

//NewClient return new Client Struct
func NewClient() *Client {
	return &Client{
		send: make(chan Message),
	}
}

func main() {
	client := NewClient()
	go client.subscribeChannels()
	go client.subscribeMessages()
	client.write()
}
