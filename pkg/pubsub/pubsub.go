package pubsub

import (
	"fmt"
	"sync"
)

// DataType is data type of message
type DataType string

// MessageChannel is a channel of pub/sub pattern
type MessageChannel struct {
	pub   chan DataType
	subs  []chan DataType
	mutex sync.Mutex
}

// Client is client of pub/sub pattern
type Client struct {
	topic map[string]*MessageChannel
}

// Pub publish message
func (m *Client) Pub(topic string, data DataType) {
	m.topic[topic].pub <- data
}

// Sub subscribe message
func (m *Client) Sub(topic string, handler func(DataType) error) {
	m.topic[topic].mutex.Lock()
	defer m.topic[topic].mutex.Unlock()
	subChannel := make(chan DataType, 10)
	m.topic[topic].subs = append(m.topic[topic].subs, subChannel)
	go func() {
		for {
			data := <-subChannel
			if err := handler(data); err != nil {
				fmt.Println(err)
			}
		}
	}()
}

// AddTopic publish message
func (m *Client) AddTopic(topic string) {
	m.topic[topic] = &MessageChannel{
		pub: make(chan DataType, 10),
	}
	go func() {
		for {
			data := <-m.topic[topic].pub
			for _, reader := range m.topic[topic].subs {
				reader <- data
			}
		}
	}()
}

// NewClient new a client
func NewClient() *Client {
	client := &Client{
		topic: make(map[string]*MessageChannel),
	}
	return client
}
