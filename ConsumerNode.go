package main

import (
	"github.com/debarshri/pipie/internals"
	"fmt"
)

func main() {

	node := pipie.Node{ProducerPort:8081, ConsumerPort:8090, ProducerHostName:"localhost", Persisted:true, DBLocation:"pipie-consumer.db"}

	consumer := pipie.CreateConsumer(node)

	consumer.Receive(func(data string){
		fmt.Println(data)
	})
}
