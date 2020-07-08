package pubsub

import (
	"fmt"
	"sync"
)

// DataType is data type of message
type DataType string

// MessageChannel is a channel of pub/sub pattern
type MessageChannel struct {
	writer  chan DataType
	readers []chan DataType
	mutex   sync.Mutex
}

// Client is client of pub/sub pattern
type Client struct {
	topic map[string]*MessageChannel
}

// Pub publish message
func (m *Client) Pub(topic string, data DataType) {
	m.topic[topic].writer <- data
}

// Sub subscribe message
func (m *Client) Sub(topic string, handler func(DataType) error) {
	m.topic[topic].mutex.Lock()
	defer m.topic[topic].mutex.Unlock()
	readChannel := make(chan DataType, 10)
	m.topic[topic].readers = append(m.topic[topic].readers, readChannel)
	go func() {
		for {
			data := <-readChannel
			if err := handler(data); err != nil {
				fmt.Println(err)
			}
		}
	}()
}

// AddTopic publish message
func (m *Client) AddTopic(topic string) {
	m.topic[topic] = &MessageChannel{
		writer: make(chan DataType, 10),
	}
	go func() {
		for {
			data := <-m.topic[topic].writer
			for _, reader := range m.topic[topic].readers {
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
