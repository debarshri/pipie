package main

import (
	"github.com/debarshri/pipie/internals"
	"time"
)

func main() {

	node := pipie.Node{ProducerPort: 8081, ConsumerPort: 8090, ConsumerHostName: "localhost", Persisted: true}

	producer := pipie.CreateServer(node)

	go producer.StartAckReceiveServer(node)

	go producer.Flush(1*time.Second)

	//This is a test

	for {
		producer.Send(time.Now().String()+
			time.Now().String()+time.Now().String()+time.Now().String()+time.Now().String()+time.Now().String())
		time.Sleep(100 * time.Millisecond)
	}
}
