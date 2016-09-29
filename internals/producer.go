package pipie

import (
	"log"
	"time"
)

func CreateServer(node Node) Producer {

	var dbLocation string
	if node.ProducerHostName == "" {

		if node.DBLocation == "" {
			dbLocation = "pipie-producer.db"
		} else {
			dbLocation = node.DBLocation
		}
		mqclient := MqClient{Hostname: node.ConsumerHostName, HostPort: node.ConsumerPort}
		return Producer{
			Server: ServerStreamAtPortWithDBLoc(node.ProducerPort,dbLocation),
			Client:mqclient,
		}
	}

	return Producer{}
}

func (p Producer) Flush(window time.Duration) {

	c := make(chan(bool))

	var lastFlush time.Time = time.Now()

	for {
		if time.Since(lastFlush) > window {
			p.Server.Flush()

			lastFlush = time.Now()

		} else {
			time.Sleep(window)
		}
	}

	c <- true
}

func (p Producer) StartAckReceiveServer(node Node) {

	c := make(chan(bool))

	p.Client.Receive(func(data string) {

		log.Println("Received ack ", data)

		p.Server.DeleteKey(data)

	})

	c <- true
}

func (p Producer) Send(payload string) {
	p.Server.Persist(payload)
}
