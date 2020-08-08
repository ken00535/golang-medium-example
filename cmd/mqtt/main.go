package main

import (
	"example/pkg/mqtt"
)

func main() {
	client := mqtt.New()
	client.Publish()
}
