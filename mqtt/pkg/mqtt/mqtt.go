package mqtt

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
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
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	clientID := strconv.Itoa(r1.Int())
	fmt.Println(clientID)
	opts := mqtt.NewClientOptions().AddBroker("tcps://127.0.0.1:1883").SetClientID(clientID)
	opts.DefaultPublishHandler = f

	tlsConfig := NewTLSConfig()
	opts.SetTLSConfig(tlsConfig)

	c.client = mqtt.NewClient(opts)
	if token := c.client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return c
}

// Publish mqtt message
func (m Client) Publish(topic string, payload interface{}) {
	text, _ := json.Marshal(payload)
	token := m.client.Publish(topic, 0, false, text)
	token.Wait()
}

// Subscribe mqtt message
func (m Client) Subscribe(topic string) {
	if token := m.client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}
