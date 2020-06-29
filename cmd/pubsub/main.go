package main

import (
	"errors"
	"example/pkg/pubsub"
	"fmt"
	"time"
)

func main() {
	client := pubsub.NewClient()
	var printMessage func(pubsub.DataType) error
	printMessage = func(msg pubsub.DataType) error {
		if msg == "error" {
			return errors.New("this is an error")
		}
		fmt.Println(msg)
		return nil
	}
	client.Sub(printMessage)
	client.Pub("Hello")
	client.Pub("error")
	time.Sleep(time.Second)
}
