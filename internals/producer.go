package pipie

import (
	"log"
)

func CreateServer(node Node) Producer {
	if node.ProducerHostName == "" {
		mqclient := MqClient{Hostname: node.ConsumerHostName, HostPort: node.ConsumerPort}
		return Producer{
			Server: ServerStreamAtPortWithDBLoc(node.ProducerPort,"pipie-producer.db"),
			Client:mqclient,
		}
	}
	return Producer{}
}

func (p Producer) StartAckReceiveServer(node Node) {

	c := make(chan(bool))

	p.Client.Receive(func(data string) {

		log.Println("Received ack ", data)

		p.Server.DeleteKey(data)

	})

	c <- true
}

func (p Producer) Send(data string) {
	p.Server.PersistedSend(data)
}
