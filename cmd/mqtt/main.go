package main

import (
	"example/pkg/mqtt"
	"time"
)

func main() {
	client := mqtt.New()
	payload := mqtt.Message{
		Header: "this is header",
	}
	client.Publish("topic/golang", payload)
	client.Subscribe("topic/golang")
	time.Sleep(30 * time.Second)
}
