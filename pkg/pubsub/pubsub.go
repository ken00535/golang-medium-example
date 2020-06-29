package pubsub

import (
	"fmt"
	"sync"
)

// DataType is data type of message
type DataType string

// Client is a client of pub/sub pattern
type Client struct {
	writer  chan DataType
	readers []chan DataType
	mutex   sync.Mutex
}

// Pub publish message
func (m *Client) Pub(data DataType) {
	m.writer <- data
}

// Sub subscribe message
func (m *Client) Sub(handler func(DataType) error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	readChannel := make(chan DataType, 10)
	m.readers = append(m.readers, readChannel)
	go func() {
		for {
			data := <-readChannel
			if err := handler(data); err != nil {
				fmt.Println(err)
			}
		}
	}()
}

// NewClient new a client
func NewClient() *Client {
	broker := &Client{
		writer: make(chan DataType, 10),
	}
	go func() {
		for {
			data := <-broker.writer
			for _, reader := range broker.readers {
				reader <- data
			}
		}
	}()
	return broker
}
