package pipie

import (
	"net"
	"github.com/boltdb/bolt"
)

type MqServer struct {
	listener  net.Listener
	queuesize int
	ttl       int64
	DB *bolt.DB
}

type Message struct {
	Key string
	Value string
}


type Consumer struct {
	Client MqClient
	Server MqServer
}

type Producer struct {
	Server MqServer
	Client MqClient
}

type MqConsumer struct {
	Hostname string
	ProducerPort int
	ConsumerPort int
}

type Node struct {
	ConsumerHostName string
	ProducerHostName string
	ProducerPort int
	ConsumerPort int
	Persisted bool
	Transient bool
	DBLocation string
}

type Database struct {
	DB *bolt.DB
}
