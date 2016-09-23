package ds

import "net"

type MqServer struct {
	listener  net.Listener
	queuesize int
	ttl       int64
}

type Message struct {
	Key string
	Value string
}

type MqClient struct {
	Hostname string
	HostPort int
}
