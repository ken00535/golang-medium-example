package main

import (
	"errors"
	"example/pkg/pubsub"
	"fmt"
	"time"
)

func main() {
	client := pubsub.NewClient()
	client.AddTopic("hello")
	client.AddTopic("echo")
	var printMessage = func(msg pubsub.DataType) error {
		if msg == "error" {
			return errors.New("This is an error")
		}
		fmt.Println(msg)
		return nil
	}
	var echoMessage = func(msg pubsub.DataType) error {
		fmt.Println(msg + " nice to meet you!")
		return nil
	}
	client.Sub("hello", printMessage)
	client.Sub("echo", echoMessage)
	client.Pub("hello", "Hello")
	client.Pub("hello", "error")
	client.Pub("echo", "Go")
	time.Sleep(time.Second)
}
