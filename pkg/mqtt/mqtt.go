package mqtt

import (
	"fmt"
	"log"
	"os"

	"github.com/eclipse/paho.mqtt.golang"
)

// Client is a mqtt client
type Client struct {
	client mqtt.Client
}

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

// New mqtt client and connect to broker
func New() Client {
	c := Client{}
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1883").SetClientID("gotrivial")

	c.client = mqtt.NewClient(opts)
	if token := c.client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	
	return c
}

// Publish mqtt message
func (m Client) Publish() {
	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		token := m.client.Publish("go-mqtt/sample", 0, false, text)
		token.Wait()
	}
}